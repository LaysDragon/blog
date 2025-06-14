package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (a *Account) ById(ctx context.Context, op perm.ResId, id int) (*domain.Account, error) {
	err := a.perm.CheckE(op, perm.ACT_READ, perm.User(id))
	if err != nil {
		return nil, err
	}
	return a.repo.ById(ctx, id)
}
