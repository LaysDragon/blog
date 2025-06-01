package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

type PostRepo interface {
	CommonRepo[PostRepo]
	ById(ctx context.Context, id int) (*domain.Post, error)
	Upsert(ctx context.Context, Post *domain.Post) (*domain.Post, error)
	Delete(ctx context.Context, id int) error
}

type Post struct {
	repo PostRepo
}

func NewPost(repo PostRepo) *Post {
	return &Post{repo}
}

func (s *Post) ById(ctx context.Context, id int) (*domain.Post, error) {
	return s.repo.ById(ctx, id)
}
