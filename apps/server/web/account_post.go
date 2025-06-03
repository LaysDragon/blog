package web

import (
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AccountCreateRequest struct {
	Email    string             `json:"email" binding:"required"`
	Username string             `json:"username" binding:"required"`
	Role     domain.AccountRole `json:"role" binding:"enum"`
	Password string             `json:"password" binding:"required"`
}

// create account
func (c *AccountController) HandlePost(ctx *gin.Context) {
	var request AccountCreateRequest = AccountCreateRequest{Role: domain.UserRole}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Status(400)
		c.log.Error("bind failed", zap.Error(err))
		return
	}

	if request.Role == domain.AdminRole {
		//TODO: verify the permission of create admin account
	}

	acc, err := c.usecase.Create(ctx, &domain.Account{
		Username: request.Username,
		Role:     request.Role,
		Email:    request.Email,
	}, request.Password)
	if err != nil {
		if err, ok := err.(usecase.ItemConflictError); ok && err.Field == "username" {
			ctx.Status(409)
		} else {
			ctx.Status(500)
		}
		c.log.Error("created account failed", zap.Error(err))
		return
	}
	ctx.JSON(200, Account{
		ID:       acc.ID,
		Username: acc.Username,
		Email:    acc.Email,
	})
}
