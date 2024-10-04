package user

import (
	user22 "Lists-app/internal/model/user"
	"context"
)

type User interface {
	Verify(ctx context.Context, user user22.User) (bool, error)
	Insert(ctx context.Context, user user22.User) error
	GetById(ctx context.Context, Id int) (user22.User, error)
	Delete(ctx context.Context, user user22.User) error
}
