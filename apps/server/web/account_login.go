package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *AccountController) HandleGet(ctx *gin.Context) {
	ctx.Status(200)
}

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

	result, err := c.usecase.Verify(ctx, request.Username, request.Password)
	if err != nil {
		c.log.Error("account verify failed", zap.Error(err))
		return
	}
	if result {
		ctx.JSON(http.StatusOK, LoginResult{
			Token: "",
		})
	} else {
		ctx.Status(http.StatusUnauthorized)
	}

}
