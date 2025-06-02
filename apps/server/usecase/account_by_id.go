package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (s *Account) ById(ctx context.Context, id int) (*domain.Account, error) {
	return s.repo.ById(ctx, id)
}
