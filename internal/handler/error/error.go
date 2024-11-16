package error

import (
	"errors"
	"net/http"

	modelerror "Lists-app/internal/model/error"
)

func Resolver(err error) int {
	switch {

	case errors.Is(err, modelerror.ErrUnknown):
		return http.StatusInternalServerError

	case errors.Is(err, modelerror.ErrWrongPassword):
		return http.StatusBadRequest

	case errors.Is(err, modelerror.UserAlreadyExistsErr):
		return http.StatusBadRequest

	case errors.Is(err, modelerror.ErrUserNotFound):
		return http.StatusBadRequest

	}

	return http.StatusInternalServerError
}
