package pgrepo

import (
	"context"
	"database/sql"
	"errors"

	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type CommonDb[T usecase.CommonRepo[T]] struct {
	db      boil.ContextExecutor
	builder func(db boil.ContextExecutor) T
}

func nilVal[T any]() T {
	var t T
	return t
}

func (c *CommonDb[T]) BeginTx(ctx context.Context) (T, error) {
	if db, ok := c.db.(*sql.DB); ok {
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return nilVal[T](), err
		}
		return c.builder(tx), nil
	}
	return nilVal[T](), errors.New("not *sql.DB")
}

func (c *CommonDb[T]) Commit() error {
	if tx, ok := c.db.(*sql.Tx); ok {
		return tx.Commit()
	}
	return errors.New("not *sql.Tx")
}

func (c *CommonDb[T]) Rollback() error {
	if tx, ok := c.db.(*sql.Tx); ok {
		return tx.Rollback()
	}
	return errors.New("not *sql.Tx")
}

func ErrorTranslate(err error) error {
	if err == sql.ErrNoRows {
		return &usecase.ItemNotExistedError{Err: err}
	}
	return err
}
