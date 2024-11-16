package user

import (
	user22 "Lists-app/internal/model/user"
	"context"
)

type User interface {
	SignIn(ctx context.Context, user user22.User) error
	SignUp(ctx context.Context, user user22.User) error
	GetById(ctx context.Context, Id int) (user22.User, error)
	Delete(ctx context.Context, user user22.User) error
}
