package usecase

import "context"

type CommonRepo[T any] interface {
	BeginTx(ctx context.Context) (T, error)
	Commit() error
	Rollback() error
}
