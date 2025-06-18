package web

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	Role string
}

type JwtHandler struct {
	secret []byte
	log    *zap.Logger
}

func NewJwtHandler(secret string, log *zap.Logger) *JwtHandler {
	return &JwtHandler{secret: []byte(secret), log: log}
}

func (j *JwtHandler) Parse(tokenStr string) (*jwt.Token, *AuthClaims, error) {
	var claim AuthClaims
	token, err := jwt.ParseWithClaims(tokenStr, &claim, func(t *jwt.Token) (interface{}, error) {
		return j.secret, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, &claim, nil
}

func (j *JwtHandler) Signed(uid int, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		AuthClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    "Server",
				Subject:   strconv.Itoa(uid),
				Audience:  jwt.ClaimStrings{},
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 3)),
				// NotBefore: jwt.NewNumericDate(time.Now()),
				// IssuedAt:  jwt.NewNumericDate(time.Now()),
				// ID: "",
			},
			Role: role,
		})
	tokenStr, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

type AuthToken struct {
	Token   *jwt.Token
	Claims  *AuthClaims
	Expired bool
}

func (j *JwtHandler) ParseMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearer := ctx.GetHeader("Authorization")
		if bearer == "" {
			return
		}
		spltBearer := strings.Split(bearer, " ")
		if len(spltBearer) != 2 {
			return
		}

		tokenStr := spltBearer[1]

		expired := false
		token, claims, err := j.Parse(tokenStr)
		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenExpired):
				expired = true
			case
				errors.Is(err, jwt.ErrTokenSignatureInvalid),
				errors.Is(err, jwt.ErrTokenMalformed):
				j.log.Error("failed to parse jwt token", zap.Error(err))
				ctx.AbortWithStatus(http.StatusBadRequest)
				return
			default:
				j.log.Error("failed to parse jwt token", zap.Error(err))
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}

		ctx.Set("token", &AuthToken{
			token,
			claims,
			expired,
		})
	}
}

func RequiredAuthMiddware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := GetToken(ctx)
		if token == nil || token.Expired || !token.Token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func GetToken(ctx context.Context) *AuthToken {
	if token, ok := ctx.Value("token").(*AuthToken); ok {
		return token
	}
	return nil

}

func GetRole(ctx context.Context) domain.AccountRole {
	token := GetToken(ctx)
	if token == nil {
		return ""
	}
	if token.Expired {
		return ""
	}

	if !token.Token.Valid {
		return ""
	}

	return (domain.AccountRole)(token.Claims.Role)
}

func GetUID(ctx context.Context) int {
	token := GetToken(ctx)
	if token == nil {
		return -1
	}

	if token.Expired {
		return -1
	}

	if !token.Token.Valid {
		return -1
	}
	uid, err := strconv.Atoi(token.Claims.Subject)
	if err != nil {
		GetLogger(ctx).Error("failed to parse jwt token uid subject", zap.Error(err))
		return -1
	}
	return uid
}
