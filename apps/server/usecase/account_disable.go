package usecase

import (
	"context"
)

func (s *Account) SetDisable(ctx context.Context, id int) error {
	return nil
	// return s.repo.Upsert(ctx, Account)
}
