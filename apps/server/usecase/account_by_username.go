package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (s *Account) ByUsername(ctx context.Context, username string) (*domain.Account, error) {
	return s.accRepo.ByUsername(ctx, username)
}
