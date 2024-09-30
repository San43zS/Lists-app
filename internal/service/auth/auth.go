package auth

import user22 "Lists-app/internal/model/user"

type Autorization interface {
	Registration(user user22.User) error
	Authenticate(user user22.User) error
	GenerateToken(user user22.User) (string, error)
}
