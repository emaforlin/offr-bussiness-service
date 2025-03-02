package service

import "errors"

var (
	ErrBusinessNotFound = errors.New("business not found")
	ErrOnBusinessDelete = errors.New("failed to delete business")
	ErrOnBusinessCreate = errors.New("failed to create business")
)
