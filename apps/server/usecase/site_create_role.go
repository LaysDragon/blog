package usecase

import (
	"context"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (s *Site) CreateRole(ctx context.Context, op perm.ResId, sid int, uid int, role domain.SiteRoleType) (bool, error) {

	//TODO: if need advanced site role management in the feature there need to add more detail
	var err error
	if role == domain.SiteOwnerRole {
		err = s.perm.CheckE(op, perm.ACT_WRITE_SITE_OWNER, perm.Site(sid))
	} else {
		err = s.perm.CheckE(op, perm.ACT_WRITE, perm.Site(sid))
	}
	if err != nil {
		return false, fmt.Errorf("create site role data failed:%w", err)
	}

	_, err = s.roleRepo.Upsert(ctx, &domain.SiteRole{
		SiteId:    sid,
		AccountId: uid,
		Role:      role,
	})
	if err != nil {
		return false, fmt.Errorf("create site role data failed:%w", err)
	}
	return true, nil
}
