package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (s *Post) Update(ctx context.Context, post *domain.Post) (*domain.Post, error) {
	return s.repo.Upsert(ctx, post)
}
