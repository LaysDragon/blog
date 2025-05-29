package service

import "github.com/LaysDragon/blog/apps/server/domain"

type PostRepository interface {
	ById(id int) (*domain.Post, error)
}

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{repo}
}

func (s *PostService) ById(id int) (*domain.Post, error) {
	return s.repo.ById(id)
}
