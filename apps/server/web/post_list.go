package web

import (
	"errors"
	"net/http"

	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PostListRequest struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"  binding:"required"`
	SiteId   int `form:"siteId"`
}

func (c *PostController) HandleList(ctx *gin.Context) {
	var request PostListRequest
	if err := ctx.ShouldBindQuery(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
		c.log.Error("bind failed", zap.Error(err))
		return
	}

	posts, err := c.usecase.List(ctx, GetUserOp(ctx), request.Page, request.PageSize, request.SiteId)
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
	ctx.JSON(200, utils.MappingFunc(posts, c.ToDto))
}
