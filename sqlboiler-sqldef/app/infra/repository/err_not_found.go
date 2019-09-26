package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/kei2100/playground-go-sql/sqlboiler-sqldef/app/domain/repository"
	perrors "github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() perrors.StackTrace
}

func errNotFound(cause error) error {
	if _, ok := cause.(stackTracer); ok {
		return fmt.Errorf("%w: %+v", repository.ErrNotFound, cause)
	}
	return fmt.Errorf("%w: %+v", repository.ErrNotFound, perrors.WithStack(cause))
}

func isSQLErrNoRows(err error) bool {
	return perrors.Cause(err) == sql.ErrNoRows || errors.Is(err, sql.ErrNoRows)
}
