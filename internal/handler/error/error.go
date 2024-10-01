package error

import (
	"errors"
	"net/http"
)

func ErrorResolver(err error) int {
	switch {

	case errors.Is(err, apperr.ErrUnknown):
		return http.StatusInternalServerError

	case errors.Is(err, apperr.ErrNoSuchCart):
		return http.StatusBadRequest

	case errors.Is(err, apperr.ErrNoSuchItem):
		return http.StatusBadRequest

	case errors.Is(err, apperr.ErrUrl):
		return http.StatusBadRequest

	case errors.Is(err, apperr.ErrItemData):
		return http.StatusBadRequest

	}

	return http.StatusInternalServerError
}
