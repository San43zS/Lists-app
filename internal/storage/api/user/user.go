package user

import (
	user2 "Lists-app/internal/model/user"
	"context"
)

type User interface {
	SignIn(ctx context.Context, user user2.User) error
	SignUp(ctx context.Context, user user2.User) error
	GetById(ctx context.Context, Id int) (user2.User, error)
	Delete(ctx context.Context, user user2.User) error
}
