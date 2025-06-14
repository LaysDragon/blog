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

func (c *AccountController) HandleList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	pageSize, err := strconv.Atoi(ctx.Query("pageSize"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	accs, err := c.usecase.List(ctx, GetUserOp(ctx), page, pageSize)
	if err != nil {
		c.log.Error("Failed to retrived account data", zap.Int("page", page), zap.Int("pageSize", pageSize), zap.Error(err))
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
	var result []Account
	for _, a := range accs {
		result = append(result, *c.ToDto(a))
	}

	ctx.JSON(http.StatusOK, result)
}
