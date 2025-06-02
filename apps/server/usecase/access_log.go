package usecase

import (
	"context"

	"github.com/LaysDragon/blog/apps/server/domain"
)

type AccessLogRepo interface {
	CommonRepo[AccessLogRepo]
	ById(ctx context.Context, id int) (*domain.AccessLog, error)
	Upsert(ctx context.Context, post *domain.AccessLog) (*domain.AccessLog, error)
	Delete(ctx context.Context, id int) error
}

type AccessLog struct {
	repo AccessLogRepo
}

func NewAccessLog(repo AccessLogRepo) *AccessLog {
	return &AccessLog{repo}
}
