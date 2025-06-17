package usecase

import (
	"context"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/utils"
)

func (s *Site) Create(ctx context.Context, op perm.ResId, Site *domain.Site) (*domain.Site, error) {
	err := s.perm.CheckE(op, perm.ACT_WRITE.Res(perm.RES_SITE), perm.System())
	if err != nil {
		return nil, fmt.Errorf("create site data failed:%w", err)
	}

	return utils.ErrorWrap(s.repo.Upsert(ctx, Site))("create site data failed:%w")
}
