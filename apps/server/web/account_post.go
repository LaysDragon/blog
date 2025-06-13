package web

import (
	"errors"
	"net/http"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
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
		ctx.Status(http.StatusBadRequest)
		c.log.Error("bind failed", zap.Error(err))
		return
	}

	acc, err := c.usecase.Create(ctx, GetUserOp(ctx), &domain.Account{
		Username: request.Username,
		Role:     request.Role,
		Email:    request.Email,
	}, request.Password)

	if err != nil {
		switch {
		case errors.Is(err, usecase.ItemConflictError{}):
			var e usecase.ItemConflictError
			if ok := errors.As(err, &e); ok && e.Field == "username" {
				ctx.Status(http.StatusConflict)
			} else {
				ctx.Status(http.StatusInternalServerError)
			}
		case errors.Is(err, perm.PermissionError{}):
			ctx.Status(http.StatusForbidden)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		c.log.Error("created account failed", zap.Error(err))
		return
	}

	ctx.JSON(200, c.ToDto(acc))
}
