package error

import (
	"errors"
	"net/http"

	modelerror "Lists-app/internal/model/error"
)

func ErrorResolver(err error) int {
	switch {

	case errors.Is(err, modelerror.ErrUnknown):
		return http.StatusInternalServerError

	case errors.Is(err, modelerror.ErrVerifyUser):
		return http.StatusBadRequest

	case errors.Is(err, modelerror.ErrUserAlreadyExists):
		return http.StatusBadRequest

	case errors.Is(err, modelerror.ErrUserNotFound):
		return http.StatusBadRequest

	}

	return http.StatusInternalServerError
}
