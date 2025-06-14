package pgrepo

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/db/pgrepo/models"
	"github.com/LaysDragon/blog/apps/server/domain"
	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AccountDb struct {
	CommonDb[usecase.AccountRepo]
}

func NewAccount(db boil.ContextExecutor) usecase.AccountRepo {
	return &AccountDb{
		CommonDb: CommonDb[usecase.AccountRepo]{
			db:      db,
			builder: NewAccount,
		},
	}
}

func (r *AccountDb) ToDb(account *domain.Account) *models.Account {
	if account == nil {
		return nil
	}
	return &models.Account{
		ID:         account.ID,
		CreatedAt:  account.CreatedAt,
		UpdatedAt:  account.UpdatedAt,
		Username:   account.Username,
		Role:       string(account.Role),
		Email:      account.Email,
		PasswdHash: account.PasswdHash,
	}
}

func (r *AccountDb) ToDomain(account *models.Account) *domain.Account {
	if account == nil {
		return nil
	}
	return &domain.Account{
		ID:         account.ID,
		CreatedAt:  account.CreatedAt,
		UpdatedAt:  account.UpdatedAt,
		Username:   account.Username,
		Role:       domain.AccountRole(account.Role),
		Email:      account.Email,
		PasswdHash: account.PasswdHash,
	}
}

func (r *AccountDb) ById(ctx context.Context, id int) (*domain.Account, error) {
	Account, err := models.FindAccount(ctx, r.db, id)

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(Account), nil
}

func (r *AccountDb) ByUsername(ctx context.Context, username string) (*domain.Account, error) {
	Account, err := models.Accounts(models.AccountWhere.Username.EQ(username)).One(ctx, r.db)

	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(Account), nil
}

type ErrorWrapped interface{ Unwrap() error }

func (r *AccountDb) Upsert(ctx context.Context, account *domain.Account) (*domain.Account, error) {
	dbAccount := r.ToDb(account)
	err := dbAccount.Upsert(ctx, r.db, true, []string{"id"}, boil.Infer(), boil.Infer())
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	return r.ToDomain(dbAccount), nil
}

func (r *AccountDb) Delete(ctx context.Context, id int) error {
	_, err := models.Accounts(models.AccountWhere.ID.EQ(id)).DeleteAll(ctx, r.db, false)
	return ErrorTranslate(err)
}

func (r *AccountDb) List(ctx context.Context, offset int, limit int) ([]*domain.Account, error) {
	accs, err := models.Accounts(Offset(offset), Limit(limit), OrderBy(models.AccountColumns.ID)).All(ctx, r.db)
	if err != nil {
		return nil, ErrorTranslate(err)
	}
	var result []*domain.Account
	for _, a := range accs {
		result = append(result, r.ToDomain(a))
	}
	return result, nil

}
