package error

import "errors"

const (
	UniqueViolationErr = "23505"
)

var (
	UserAlreadyExistsErr = errors.New("user already exists")

	ErrUserNotFound = errors.New("user not found")
	ErrUnknown      = errors.New("something went wrong")

	ErrWrongEmail    = errors.New("email  is incorrect")
	ErrWrongPassword = errors.New("wrong password")
)
