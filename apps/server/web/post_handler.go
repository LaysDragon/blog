package web

import (
	"database/sql"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	usecase *usecase.PostUseCase
}

func NewPostController(usecase *usecase.PostUseCase) *PostController {
	return &PostController{usecase}
}

func (c *PostController) HandleGetPost(ctx *gin.Context) {
	post, err := c.usecase.ById(0)
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
