package web

import (
	"net/http"
	"strconv"

	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *AccountController) HandleGet(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.Status(http.StatusBadRequest)
	}

	//TODO: add permission check

	acc, err := c.usecase.ById(ctx, id)
	if err != nil {
		if _, ok := err.(usecase.ItemNotExistedError); ok {
			ctx.Status(http.StatusNotFound)
			return
		}
		c.log.Error("Failed to retrived account data", zap.Int("id", id), zap.Error(err))
		ctx.Status(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, c.ToDto(acc))
}
