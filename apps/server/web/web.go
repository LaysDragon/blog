package web

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type WebConfig interface {
	GetJwtSecret() string
}

var Module = fx.Module("web",
	fx.Provide(
		GetValidator,
		func(config WebConfig, log *zap.Logger) *JwtHandler {
			return NewJwtHandler(config.GetJwtSecret(), log)
		},
		NewPostController,
		NewAccountController,
		NewSiteController,
	),

	fx.Invoke(
		SetupValidation,
		SetupRouter,
	))

const LOGGER_KEY = "APP_WIDE_LOGGER"

func RegisterLogger(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, LOGGER_KEY, log)
}

func RegisterGinLogger(ctx *gin.Context, log *zap.Logger) {
	ctx.Set(LOGGER_KEY, log)
}

func GetLogger(ctx context.Context) *zap.Logger {
	if v, ok := ctx.Value(LOGGER_KEY).(*zap.Logger); ok {
		return v
	}
	return nil
}

func CtxLoggerMidleware(log *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		RegisterGinLogger(ctx, log)
	}
}

func GetUserOp(ctx context.Context) perm.ResId {
	return perm.User(GetUID(ctx))
}
