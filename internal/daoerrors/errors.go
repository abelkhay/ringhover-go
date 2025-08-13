package daoerrors

import "errors"

var (
	ErrNotFound = errors.New("Not Found")
	ErrBadInput = errors.New("Bad Input")
	ErrConflict = errors.New("Conflict")
)
