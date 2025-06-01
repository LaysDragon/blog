package web

import (
	"database/sql"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
	"log"
)

type PostController struct {
	usecase *usecase.Post
}

func NewPostController(usecase *usecase.Post) *PostController {
	return &PostController{usecase}
}

func (c *PostController) HandleGetPost(ctx *gin.Context) {
	post, err := c.usecase.ById(ctx, 1)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(404)
		} else {
			ctx.String(500, err.Error())
		}
		log.Println(err)
		return
	}
	ctx.JSON(200, post)
}
