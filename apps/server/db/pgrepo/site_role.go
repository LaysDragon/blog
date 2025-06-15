package pgrepo

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SiteRoleDb struct {
	CommonDb[usecase.SiteRoleRepo]
}

func NewSiteRole(db stdlibTransactor.DBGetter) usecase.SiteRoleRepo {
	return &SiteRoleDb{
		CommonDb: CommonDb[usecase.SiteRoleRepo]{
			db:      db,
			builder: NewSiteRole,
		},
	}
}

func (r *SiteRoleDb) ToDb(role *domain.SiteRole) *models.SiteRole {
	if role == nil {
		return nil
	}
	return &models.SiteRole{
		AccountID: role.AccountId,
		SiteID:    role.SiteId,
		Role:      string(role.Role),
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (r *SiteRoleDb) ToDomain(role *models.SiteRole) *domain.SiteRole {
	if role == nil {
		return nil
	}
	return &domain.SiteRole{
		AccountId: role.AccountID,
		SiteId:    role.SiteID,
		Role:      domain.SiteRoleType(role.Role),
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

func (r *SiteRoleDb) ById(ctx context.Context, sid int, uid int) (*domain.SiteRole, error) {
	SiteRole, err := models.FindSiteRole(ctx, r.db(ctx), uid, sid)

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(SiteRole), nil
}

func (r *SiteRoleDb) ByUid(ctx context.Context, id int) ([]*domain.SiteRole, error) {
	roles, err := models.SiteRoles(models.SiteRoleWhere.AccountID.EQ(id)).All(ctx, r.db(ctx))

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return mappingFunc(roles, r.ToDomain), nil
}

func (r *SiteRoleDb) BySid(ctx context.Context, id int) ([]*domain.SiteRole, error) {
	roles, err := models.SiteRoles(models.SiteRoleWhere.SiteID.EQ(id)).All(ctx, r.db(ctx))

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return mappingFunc(roles, r.ToDomain), nil
}

func (r *SiteRoleDb) Upsert(ctx context.Context, role *domain.SiteRole) (*domain.SiteRole, error) {
	dbSiteRole := r.ToDb(role)
	err := dbSiteRole.Upsert(ctx, r.db(ctx), true, []string{"site_id", "account_id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(dbSiteRole), nil
}

func (r *SiteRoleDb) Delete(ctx context.Context, sid int, uid int) error {
	_, err := models.SiteRoles(models.SiteRoleWhere.SiteID.EQ(sid), models.SiteRoleWhere.AccountID.EQ(uid)).DeleteAll(ctx, r.db(ctx))
	return ErrorTranslate(err)
}

// TODO seems useless without filter option
func (r *SiteRoleDb) List(ctx context.Context, offset int, limit int) ([]*domain.SiteRole, error) {
	accs, err := models.SiteRoles(Offset(offset), Limit(limit), OrderBy(models.SiteRoleColumns.SiteID)).All(ctx, r.db(ctx))
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return mappingFunc(accs, r.ToDomain), nil

}
