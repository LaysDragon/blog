package web

import "go.uber.org/zap"

func logObj(log *zap.Logger, obj any) {
	log.Debug("check obj content", zap.Any("obj", obj))
}
