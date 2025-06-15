package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

type SiteRepo interface {
	CommonRepo[SiteRepo]
	List(ctx context.Context, offset int, limit int) ([]*domain.Site, error)
	ById(ctx context.Context, id int) (*domain.Site, error)
	Upsert(ctx context.Context, post *domain.Site) (*domain.Site, error)
	Delete(ctx context.Context, id int) error
}

type Site struct {
	repo SiteRepo
}

func NewSite(repo SiteRepo) *Site {
	return &Site{repo}
}
