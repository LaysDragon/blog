package web

import (
	"errors"
	"net/http"

	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
)

func (c *PostController) HandleGet(ctx *gin.Context) {
	post, err := c.usecase.ById(ctx, 1)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ItemNotExistedError{}):
			ctx.Status(http.StatusNotFound)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		c.log.Error(err)
		return
	}
	ctx.JSON(200, post)
}
