package pgrepo

import (
	"database/sql"
	"errors"
	"regexp"

	usecase "github.com/LaysDragon/blog/apps/server/usecase"
	stdlibTransactor "github.com/Thiht/transactor/stdlib"
	"github.com/lib/pq"
	"go.uber.org/fx"
)

var Module = fx.Module("db",
	fx.Provide(
		NewPost,
		NewAccount,
		NewSite,
		NewSiteRole,
	),
)

type CommonDb[T usecase.CommonRepo[T]] struct {
	db      stdlibTransactor.DBGetter
	builder func(db stdlibTransactor.DBGetter) T
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
