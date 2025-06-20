package pgrepo

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/LaysDragon/blog/apps/server/utils"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SiteDb struct {
	CommonDb[usecase.SiteRepo]
}

func NewSite(db stdlibTransactor.DBGetter) usecase.SiteRepo {
	return &SiteDb{
		CommonDb: CommonDb[usecase.SiteRepo]{
			db:      db,
			builder: NewSite,
		},
	}
}

func (r *SiteDb) ToDb(site *domain.Site) *models.Site {
	if site == nil {
		return nil
	}
	return &models.Site{
		ID:        site.ID,
		Name:      site.Name,
		CreatedAt: site.CreatedAt,
		UpdatedAt: site.UpdatedAt,
	}
}

func (r *SiteDb) ToDomain(site *models.Site) *domain.Site {
	if site == nil {
		return nil
	}
	return &domain.Site{
		ID:        site.ID,
		Name:      site.Name,
		CreatedAt: site.CreatedAt,
		UpdatedAt: site.UpdatedAt,
	}
}

func (r *SiteDb) ById(ctx context.Context, id int) (*domain.Site, error) {
	site, err := models.FindSite(ctx, r.db(ctx), id)

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(site), nil
}

func (r *SiteDb) Upsert(ctx context.Context, site *domain.Site) (*domain.Site, error) {
	dbSite := r.ToDb(site)
	err := dbSite.Upsert(ctx, r.db(ctx), true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(dbSite), nil
}

func (r *SiteDb) Delete(ctx context.Context, id int) error {
	affRow, err := models.Sites(models.SiteWhere.ID.EQ(id)).DeleteAll(ctx, r.db(ctx), false)
	if affRow == 0 {
		return usecase.ItemNotExistedError{}
	}
	return ErrorTranslate(err)
}

func (r *SiteDb) List(ctx context.Context, offset int, limit int, uid int) ([]*domain.Site, error) {
	if uid > 0 {
		sitesRoles, err := models.SiteRoles(Load(models.SiteRoleRels.Site), Offset(offset), Limit(limit), OrderBy(models.SiteRoleColumns.SiteID), models.SiteRoleWhere.AccountID.EQ(uid)).All(ctx, r.db(ctx))
		if err != nil {
			return nil, ErrorTranslate(err)
		}
		return utils.MappingFunc(sitesRoles, func(site *models.SiteRole) *domain.Site {
			return r.ToDomain(site.R.Site)
		}), nil
	} else {
		sites, err := models.Sites(Offset(offset), Limit(limit), OrderBy(models.SiteColumns.ID)).All(ctx, r.db(ctx))
		if err != nil {
			return nil, ErrorTranslate(err)
		}

		return utils.MappingFunc(sites, r.ToDomain), nil
	}

}
