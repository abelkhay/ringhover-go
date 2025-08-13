package httperr

import (
	"errors"
	"net/http"

	"ringhover-go/internal/daoerrors"
)

func StatusCode(err error) int {
	switch {
	case errors.Is(err, daoerrors.ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, daoerrors.ErrBadInput):
		return http.StatusBadRequest
	case errors.Is(err, daoerrors.ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func PublicMessage(err error) string {
	switch {
	case errors.Is(err, daoerrors.ErrNotFound):
		return "resource not found"
	case errors.Is(err, daoerrors.ErrBadInput):
		return "invalid request"
	case errors.Is(err, daoerrors.ErrConflict):
		return "conflict"
	default:
		return "internal error"
	}
}
