package usecase

import (
	"context"
	"errors"

	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/matthewhartstonge/argon2"
)

func (a *Account) Verify(ctx context.Context, username string, password string) (bool, *domain.Account, error) {
	account, err := a.ByUsername(ctx, username)
	if err != nil {
		switch {
		case errors.Is(err, ItemNotExistedError{}):
			return false, nil, nil
		default:
			return false, nil, err
		}
	}
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(account.PasswdHash))
	if err != nil {
		return false, nil, err
	}
	return ok, account, nil
}
