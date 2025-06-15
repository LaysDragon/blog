package usecase

import (
	"fmt"
)

type CommonRepo[T any] interface {
	// BeginTx(ctx context.Context) (T, error)
	// Commit() error
	// Rollback() error
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

func (e ItemNotExistedError) Is(err error) bool {
	_, ok := err.(ItemNotExistedError)
	return ok
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

func (e ItemConflictError) Is(err error) bool {
	_, ok := err.(ItemConflictError)
	return ok
}

func errorWrap[T any](val T, err error) func(string) (T, error) {
	return func(msg string) (T, error) {
		if err != nil {
			return val, fmt.Errorf(msg, err)
		}
		return val, err
	}
}
