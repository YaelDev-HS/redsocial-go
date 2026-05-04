package data

import "errors"

var (
	ErrDuplicatedKey = errors.New("err duplicated key")
	DuplicatedKey    = "duplicate key value violates unique constraint"
	ErrModelNotFound = errors.New("model not found")
)
