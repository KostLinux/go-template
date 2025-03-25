package constants

import "errors"

var (
	ErrNotFound = errors.New("item not found")
	ErrConflict = errors.New("item already exists")
)
