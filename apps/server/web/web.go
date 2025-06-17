package web

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/perm"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func logObj(log *zap.Logger, obj any) {
	log.Debug("check obj content", zap.Any("obj", obj))
}

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

// TODO: move to utils
func mappingFunc[S any, T any](source []S, mapper func(S) T) []T {
	result := make([]T, 0) //prevent nil slice turn into null json
	for _, s := range source {
		result = append(result, mapper(s))
	}
	return result

}
