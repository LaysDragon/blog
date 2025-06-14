package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *AccountController) HandleGet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	acc, err := c.usecase.ById(ctx, GetUserOp(ctx), id)
	if err != nil {
		c.log.Error("Failed to retrived account data", zap.Int("id", id), zap.Error(err))
		switch {
		case errors.Is(err, usecase.ItemNotExistedError{}):
			ctx.Status(http.StatusNotFound)
		case errors.Is(err, perm.PermissionError{}):
			ctx.Status(http.StatusForbidden)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		return
	}
	ctx.JSON(http.StatusOK, c.ToDto(acc))
}
