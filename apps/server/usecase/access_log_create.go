package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (s *AccessLog) Create(ctx context.Context, uid int, method string) (*domain.AccessLog, error) {
	return s.repo.Upsert(ctx, &domain.AccessLog{
		UserID: uid,
		Method: method,
	})
}
