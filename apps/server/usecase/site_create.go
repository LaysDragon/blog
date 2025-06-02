package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (s *Site) Create(ctx context.Context, Site *domain.Site) (*domain.Site, error) {
	return s.repo.Upsert(ctx, Site)
}
