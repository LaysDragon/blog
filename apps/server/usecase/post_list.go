package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (s *Post) List(ctx context.Context, op perm.ResId, page int, pageSize int, sid int) ([]*domain.Post, error) {
	res := perm.Site(sid)
	if sid <= 0 {
		res = perm.SiteWild()
	}
	err := s.perm.CheckE(op, perm.ACT_LIST.Res(perm.RES_POST), res)
	if err != nil {
		return nil, err
	}
	return s.repo.List(ctx, pageSize*page, pageSize, sid)
}
