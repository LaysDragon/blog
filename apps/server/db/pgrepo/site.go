package pgrepo

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SiteDb struct {
	CommonDb[usecase.SiteRepo]
}

func NewSite(db boil.ContextExecutor) usecase.SiteRepo {
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
	Site, err := models.FindSite(ctx, r.db, id)

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(Site), nil
}

func (r *SiteDb) Upsert(ctx context.Context, site *domain.Site) (*domain.Site, error) {
	dbSite := r.ToDb(site)
	err := dbSite.Upsert(ctx, r.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(dbSite), nil
}

func (r *SiteDb) Delete(ctx context.Context, id int) error {
	_, err := models.Sites(models.SiteWhere.ID.EQ(id)).DeleteAll(ctx, r.db, false)
	return ErrorTranslate(err)
}

func (r *SiteDb) List(ctx context.Context, offset int, limit int) ([]*domain.Site, error) {
	accs, err := models.Sites(Offset(offset), Limit(limit), OrderBy(models.SiteColumns.ID)).All(ctx, r.db)
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	var result []*domain.Site
	for _, a := range accs {
		result = append(result, r.ToDomain(a))
	}
	return result, nil

}
