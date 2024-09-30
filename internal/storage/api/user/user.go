package user

import (
	user2 "Lists-app/internal/model/user"
)

type User interface {
	Verify(user user2.User) (bool, error)
	Insert(user user2.User) error
	GetById(Id int) (user2.User, error)
	Delete(user user2.User) error
}
