package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

func (a *Account) Create(ctx context.Context, account *domain.Account, password string) (*domain.Account, error) {
	encoded, err := a.argon2.HashEncoded([]byte(password))
	if err != nil {
		return nil, err
	}
	account.PasswdHash = string(encoded)
	return a.repo.Upsert(ctx, account)
}
