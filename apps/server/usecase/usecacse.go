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
	return fmt.Sprintf("Item Not Existed: %v", e.Err)
}

func (e ItemNotExistedError) Unwrap() error {
	return e.Err
}

type ItemConflictError struct {
	Err   error
	Field string
}

func (e ItemConflictError) Error() string {
	return fmt.Sprintf("Item Conflict on (%v): %v", e.Field, e.Err)
}

func (e ItemConflictError) Unwrap() error {
	return e.Err
}
