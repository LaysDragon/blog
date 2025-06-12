package web

import (
	"github.com/LaysDragon/blog/apps/server/domain"
	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/LaysDragon/blog/apps/server/usecase"
	"go.uber.org/zap"
)

// Account DTO
type Account struct {
	ID       int
	Username string
	Email    string
	// Role     string
}

type AccountController struct {
	usecase *usecase.Account
	jwt     *JwtHandler
	log     *zap.Logger
	perm    *perm.Perm
}

func NewAccountController(usecase *usecase.Account, log *zap.Logger, perm *perm.Perm, jwt *JwtHandler) *AccountController {
	return &AccountController{usecase, jwt, log, perm}
}

func (c *AccountController) ToDto(acc *domain.Account) *Account {
	return &Account{
		ID:       acc.ID,
		Username: acc.Username,
		Email:    acc.Email,
	}

}
