package pgrepo

import (
	"database/sql"
	"errors"
	"regexp"

	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/lib/pq"
)

type CommonDb[T usecase.CommonRepo[T]] struct {
	db      stdlibTransactor.DBGetter
	builder func(db stdlibTransactor.DBGetter) T
}

func nilVal[T any]() T {
	var t T
	return t
}

// func (c *CommonDb[T]) BeginTx(ctx context.Context) (T, error) {
// 	// if db, ok := c.db.(*sql.DB); ok {
// 	// 	tx, err := db.BeginTx(ctx, nil)
// 	// 	if err != nil {
// 	// 		return nilVal[T](), err
// 	// 	}
// 	// 	return c.builder(tx), nil
// 	// }
// 	return nilVal[T](), errors.New("not *sql.DB")
// }

// func (c *CommonDb[T]) Commit() error {
// 	// if tx, ok := c.db.(*sql.Tx); ok {
// 	// 	return tx.Commit()
// 	// }
// 	return errors.New("not *sql.Tx")
// }

// func (c *CommonDb[T]) Rollback() error {
// 	// if tx, ok := c.db.(*sql.Tx); ok {
// 	// 	return tx.Rollback()
// 	// }
// 	return errors.New("not *sql.Tx")
// }

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

// TODO: move to utils
type ErrorWrapped interface{ Unwrap() error }

// TODO: move to utils
func mappingFunc[S any, T any](source []S, mapper func(S) T) []T {
	result := make([]T, 0)
	for _, s := range source {
		result = append(result, mapper(s))
	}
	return result

}
