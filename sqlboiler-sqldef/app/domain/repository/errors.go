package repository

import (
	"errors"

	"golang.org/x/xerrors"
)

// IsNotFound reports whether the err is model not found error
func IsNotFound(err error) bool {
	return errors.Is(err, ErrNotFound)
}

// ErrNotFound represents `model not found` error.
// Comparision to ErrNotFound should be `IsErrNotFound(err)`, instead of `err == ErrNotFound`
var ErrNotFound = errors.New("model not found")

// IsErrNotFound reports whether the err is `model not found` error
func IsErrNotFound(err error) bool {
	return xerrors.Is(err, ErrNotFound)
}
