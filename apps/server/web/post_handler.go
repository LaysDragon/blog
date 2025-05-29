package web

import (
	"database/sql"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	service *service.PostService
}

func NewPostController(service *service.PostService) *PostController {
	return &PostController{service}
}

func (c *PostController) HandleGetPost(ctx *gin.Context) {
	post, err := c.service.ById(0)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(404)
		} else {
			ctx.String(500, err.Error())
		}
		fmt.Println(err)
		return
	}
	ctx.JSON(200, post)
}
