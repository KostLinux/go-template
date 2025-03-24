package constants

import "errors"

var (
	ErrNotFound = errors.New("not found")
	ErrConflict = errors.New("already exists")
)
