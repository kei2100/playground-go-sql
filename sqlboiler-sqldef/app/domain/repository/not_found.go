package repository

import (
	"database/sql"

	"github.com/pkg/errors"
)

// IsNotFound reports whether the err is model not found error
func IsNotFound(err error) bool {
	// TODO go 1.13 errors
	if err := errors.Cause(err); err == sql.ErrNoRows {
		return true
	}
	return false
}
