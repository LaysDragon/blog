package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/perm"
)

func (s *Post) Delete(ctx context.Context, op perm.ResId, id int) error {

	err := s.perm.CheckE(op, perm.ACT_DELETE.Res(perm.RES_POST), perm.Post(id))
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, id)
}
