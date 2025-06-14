package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (a *Account) List(ctx context.Context, op perm.ResId, page int, pageSize int) ([]*domain.Account, error) {
	err := a.perm.CheckE(op, perm.ACT_LIST.Res(perm.RES_USER), perm.System())
	if err != nil {
		return nil, err
	}
	return a.repo.List(ctx, pageSize*page, pageSize)
}
