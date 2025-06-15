package usecase

import (
	"context"
	"fmt"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
)

func (a *Account) Create(ctx context.Context, op perm.ResId, account *domain.Account, password string) (*domain.Account, error) {
	if account.Role == domain.AdminRole {
		err := a.perm.CheckE(op, perm.ACT_WRITE_USER_ADMIN, perm.System())
		if err != nil {
			return nil, fmt.Errorf("create account failed:%w", err)
		}
	}

	encoded, err := a.argon2.HashEncoded([]byte(password))
	if err != nil {
		return nil, fmt.Errorf("create account failed:%w", err)
	}
	account.PasswdHash = string(encoded)
	account, err = a.accRepo.Upsert(ctx, account)
	if err != nil {
		return nil, fmt.Errorf("create account data failed:%w", err)
	}

	site, err := a.siteUse.Create(ctx, perm.UserSystem(op), &domain.Site{
		Name: fmt.Sprintf("User%v's site", account.ID),
	})
	if err != nil {
		return nil, err
	}

	_, err = a.siteUse.CreateRole(ctx, perm.UserSystem(op), site.ID, account.ID, domain.SiteOwnerRole)
	if err != nil {
		return nil, err
	}

	err = a.perm.Logic.AddAccount(account)
	if err != nil {
		return nil, err
	}
	err = a.perm.Logic.AddSite(site, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
