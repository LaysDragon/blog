package web

import (
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"go.uber.org/zap"
)

type PostController struct {
	usecase *usecase.Post
	log     *zap.SugaredLogger
}

func NewPostController(usecase *usecase.Post, log *zap.SugaredLogger) *PostController {
	return &PostController{usecase, log}
}

// Post DTO
type Post struct {
	Id      int
	Content string
}

func (c *PostController) ToDto(post *domain.Post) *Post {
	return &Post{
		Id:      post.ID,
		Content: post.Content,
	}

}
