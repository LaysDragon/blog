package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/Thiht/transactor"
	"github.com/matthewhartstonge/argon2"
)

type AccountRepo interface {
	CommonRepo[AccountRepo]
	List(ctx context.Context, offset int, limit int) ([]*domain.Account, error)
	ById(ctx context.Context, id int) (*domain.Account, error)
	ByUsername(ctx context.Context, username string) (*domain.Account, error)
	Upsert(ctx context.Context, post *domain.Account) (*domain.Account, error)
	Delete(ctx context.Context, id int) error
}

type Account struct {
	accRepo    AccountRepo
	siteUse    *Site
	argon2     argon2.Config
	transactor transactor.Transactor
	perm       *perm.Perm
}

func NewAccount(accRepo AccountRepo, siteUse *Site, transactor transactor.Transactor, perm *perm.Perm) *Account {
	argon := argon2.Config{
		HashLength:  32,
		SaltLength:  16,
		TimeCost:    1,
		MemoryCost:  16 * 1024,
		Parallelism: 4,
		Mode:        argon2.ModeArgon2id,
		Version:     argon2.Version13,
	}
	return &Account{accRepo, siteUse, argon, transactor, perm}
}
