package usecase

import (
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/service"
)

type PostRepository interface {
	ById(id int) (*domain.Post, error)
}

type PostUseCase struct {
	repo    PostRepository
	service *service.PostService
}

func NewPostUseCase(repo PostRepository, service *service.PostService) *PostUseCase {
	return &PostUseCase{repo, service}
}

func (s *PostUseCase) ById(id int) (*domain.Post, error) {
	return s.repo.ById(id)
}
