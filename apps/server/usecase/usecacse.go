package usecase

import (
	"context"
	"fmt"
)

type CommonRepo[T any] interface {
	BeginTx(ctx context.Context) (T, error)
	Commit() error
	Rollback() error
}

type ItemNotExistedError struct {
	Err error
}

func (e ItemNotExistedError) Error() string {
	return fmt.Sprintf("Item Not Existed : err %v", e.Err)
}

func (e ItemNotExistedError) Unwrap() error {
	return e.Err
}
