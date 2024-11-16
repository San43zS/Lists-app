package user

import (
	user2 "Lists-app/internal/model/user"
	"Lists-app/internal/service/api/user"
	store "Lists-app/internal/storage/api/user"
	"context"
)

type service struct {
	storage store.User
}

func New(storage user.User) user.User {
	return &service{
		storage: storage,
	}
}

func (s service) SignIn(ctx context.Context, user user2.User) error {
	return s.storage.SignIn(ctx, user)
}

func (s service) SignUp(ctx context.Context, user user2.User) error {
	return s.storage.SignUp(ctx, user)
}

func (s service) GetById(ctx context.Context, Id int) (user2.User, error) {
	return s.storage.GetById(ctx, Id)
}

func (s service) Delete(ctx context.Context, user user2.User) error {
	return s.storage.Delete(ctx, user)
}
