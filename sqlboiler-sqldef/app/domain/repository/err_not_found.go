package repository

import (
	"errors"
)

// ErrNotFound represents model not found error.
// Comparision to ErrNotFound should be `IsNotFound(err)`, instead of `err == ErrNotFound`
var ErrNotFound = errors.New("repository: model not found")

// IsNotFound reports whether the err is model not found error
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}
