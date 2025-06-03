package web

import (
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
}

func NewAccountController(usecase *usecase.Account, log *zap.Logger, jwt *JwtHandler) *AccountController {
	return &AccountController{usecase, jwt, log}
}
