package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResult struct {
	Token string `json:"token"`
}

// create account
func (c *AccountController) HandleLogin(ctx *gin.Context) {
	var request AccountLoginRequest = AccountLoginRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
		c.log.Error("bind failed", zap.Error(err))
		return
	}

	result, account, err := c.usecase.Verify(ctx, request.Username, request.Password)
	if err != nil {
		c.log.Error("account verify failed", zap.Error(err))
		ctx.Status(http.StatusInternalServerError)
		return
	}
	if result {
		token, err := c.jwt.Signed(account.ID, string(account.Role))
		if err != nil {
			c.log.Error("jwt signing failed", zap.Error(err))
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, LoginResult{
			Token: token,
		})
	} else {
		ctx.Status(http.StatusUnauthorized)
	}

}
