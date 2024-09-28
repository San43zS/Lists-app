package user

import (
	user2 "Lists-app/internal/model/user"
)

type User interface {
	Verify(user user2.User) error
	Insert(user user2.User) error
}
