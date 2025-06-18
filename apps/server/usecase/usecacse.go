package usecase

import (
	"fmt"

	"go.uber.org/fx"
)

var Module = fx.Module("usecase",
	fx.Provide(
		NewPost,
		NewSite,
		NewAccount,
	),
)

type CommonRepo[T any] interface {
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
