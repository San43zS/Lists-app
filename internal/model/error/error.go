package error

import "errors"

const (
	UniqueViolationErr = "23505"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")

	ErrUserNotFound = errors.New("user not found")
	ErrUnknown      = errors.New("something went wrong")

	ErrVerifyUser = errors.New("Email or password or username is incorrect")
)
