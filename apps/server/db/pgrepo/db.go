package pgrepo

import (
	"context"
	"database/sql"
	"errors"
	"regexp"

	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	"github.com/lib/pq"
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

func unpackPqError(err error) *pq.Error {
	pqError := new(pq.Error)
	if errors.As(err, &pqError) {
		return pqError
	}
	return nil
}

var uniqueViolationPattern = regexp.MustCompile(`Key \((\w+)\)=\(\w*\) already exists.`)

func extractUniqueViolationField(msg string) string {
	found := uniqueViolationPattern.FindStringSubmatch(msg)
	if len(found) == 0 {
		return ""
	}
	return found[1]
}

func ErrorTranslate(err error) error {
	if err == sql.ErrNoRows {
		return &usecase.ItemNotExistedError{Err: err}
	}
	if err := unpackPqError(err); err != nil {
		if err.Code.Name() == "unique_violation" {
			return usecase.ItemConflictError{Field: extractUniqueViolationField(err.Detail), Err: err}
		}
		return err
	}
	return err
}

type ErrorWrapped interface{ Unwrap() error }

func mappingFunc[S any, T any](source []S, mapper func(S) T) []T {
	var result []T
	for _, s := range source {
		result = append(result, mapper(s))
	}
	return result

}
