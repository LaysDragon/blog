package web

import (
	"database/sql"

	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PostController struct {
	usecase *usecase.Post
	log     *zap.SugaredLogger
}

func NewPostController(usecase *usecase.Post, log *zap.SugaredLogger) *PostController {
	return &PostController{usecase, log}
}

func (c *PostController) HandleGet(ctx *gin.Context) {
	post, err := c.usecase.ById(ctx, 1)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(404)
		} else {
			ctx.String(500, err.Error())
		}
		c.log.Error(err)
		return
	}
	ctx.JSON(200, post)
}
