package auth

import (
	user22 "Lists-app/internal/model/user"
	user222 "Lists-app/internal/storage/api/user"
)

type Autorization interface {
	Registration() user222.User
	Authenticate() user22.User
	GenerateToken() user22.User
}
