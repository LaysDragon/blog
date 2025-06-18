package web

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
)

func (c *PostController) HandleDelete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	err = c.usecase.Delete(ctx, GetUserOp(ctx), id)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ItemNotExistedError{}):
			ctx.Status(http.StatusNotFound)
		case errors.Is(err, perm.PermissionError{}):
			ctx.Status(http.StatusForbidden)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		c.log.Error(err)
		return
	}
	ctx.Status(200)
}
