package model

import (
	"errors"
)

var (
	ErrNotPersisted = errors.New("record has not persisted")
	ErrInvalidID    = errors.New("invalid BSON object id")
	ErrInvalidArgs  = errors.New("invalid arguments of the query method")
)
