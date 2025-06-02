package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (s *Post) ById(ctx context.Context, id int) (*domain.Post, error) {
	return s.repo.ById(ctx, id)
}
