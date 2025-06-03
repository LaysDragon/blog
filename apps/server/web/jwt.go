package web

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	jwt.RegisteredClaims
	Role string
}

func NewJwtHandler(secret string) *JwtHandler {
	return &JwtHandler{secret: []byte(secret)}
}

type JwtHandler struct {
	secret []byte
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
