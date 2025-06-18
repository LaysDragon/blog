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

type PostCreateRequest struct {
	Content string `json:"content" binding:"required"`
	SiteId  int    `json:"site_id" binding:"required"`
}

func (c *PostController) HandlePost(ctx *gin.Context) {
	var request PostCreateRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Status(http.StatusBadRequest)
		c.log.Error("bind failed", zap.Error(err))
		return
	}

	post, err := c.usecase.Create(ctx, GetUserOp(ctx), &domain.Post{
		Content: request.Content,
		SiteID:  request.SiteId,
	})

	if err != nil {
		switch {
		case errors.Is(err, usecase.ItemConflictError{}):
			ctx.Status(http.StatusConflict)
		case errors.Is(err, perm.PermissionError{}):
			ctx.Status(http.StatusForbidden)
		default:
			ctx.Status(http.StatusInternalServerError)
		}
		c.log.Error("created account failed", zap.Error(err))
		return
	}

	ctx.JSON(200, c.ToDto(post))
}
