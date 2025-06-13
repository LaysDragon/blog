package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (a *Account) Create(ctx context.Context, op perm.ResId, account *domain.Account, password string) (*domain.Account, error) {
	if account.Role == domain.AdminRole {
		err := a.perm.CheckE(op, perm.ACT_WRITE_USER_ADMIN, perm.System())
		if err != nil {
			return nil, err
		}
	}

	encoded, err := a.argon2.HashEncoded([]byte(password))
	if err != nil {
		return nil, err
	}
	account.PasswdHash = string(encoded)
	return a.repo.Upsert(ctx, account)
}
