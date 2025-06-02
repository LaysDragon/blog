package usecase

import "context"
import "github.com/matthewhartstonge/argon2"

func (a *Account) Verify(ctx context.Context, username string, password string) (bool, error) {
	account, err := a.ByUsername(ctx, username)
	if err != nil {
		return false, err
	}
	ok, err := argon2.VerifyEncoded([]byte(password), []byte(account.PasswdHash))
	if err != nil {
		return false, err
	}
	return ok, nil
}
