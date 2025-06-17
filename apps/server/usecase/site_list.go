package usecase

import (
	"context"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (s *Site) List(ctx context.Context, op perm.ResId, offset int, limit int, uid int) ([]*domain.Site, error) {

	res := perm.User(uid)
	if uid <= 0 {
		res = perm.System()
	}
	err := s.perm.CheckE(op, perm.ACT_LIST.Res(perm.RES_SITE), res)
	if err != nil {
		return nil, fmt.Errorf("query site data failed:%w", err)
	}

	return errorWrap(s.repo.List(ctx, offset, limit, uid))("create site data failed:%w")
}
