package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

type SiteRepo interface {
	CommonRepo[SiteRepo]
	List(ctx context.Context, offset int, limit int, uid int) ([]*domain.Site, error)
	ById(ctx context.Context, id int) (*domain.Site, error)
	Upsert(ctx context.Context, site *domain.Site) (*domain.Site, error)
	Delete(ctx context.Context, id int) error
}

type SiteRoleRepo interface {
	CommonRepo[SiteRoleRepo]
	List(ctx context.Context, offset int, limit int) ([]*domain.SiteRole, error)
	ById(ctx context.Context, sid int, uid int) (*domain.SiteRole, error)
	ByUid(ctx context.Context, uid int) ([]*domain.SiteRole, error)
	BySid(ctx context.Context, sid int) ([]*domain.SiteRole, error)
	Upsert(ctx context.Context, role *domain.SiteRole) (*domain.SiteRole, error)
	Delete(ctx context.Context, sid int, uid int) error
}

type Site struct {
	repo     SiteRepo
	roleRepo SiteRoleRepo
	perm     *perm.Perm
}

func NewSite(repo SiteRepo, roleRepo SiteRoleRepo, perm *perm.Perm) *Site {
	return &Site{repo, roleRepo, perm}
}
