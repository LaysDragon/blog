package service

import "github.com/LaysDragon/blog/apps/server/domain"

type PostRepository interface {
	ById(id int) (*domain.Post, error)
}

type PostService struct{}

func NewPostService() *PostService {
	return &PostService{}
}
