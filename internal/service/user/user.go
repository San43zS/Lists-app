package user

import (
	service2 "Lists-app/internal/service"
	"Lists-app/internal/service/api/user"
	store "Lists-app/internal/storage/api/user"
	"context"
)

type service struct {
	storage *user.User
}

func New(storage store.User) service2.Service {

}

func (s service) Verify(ctx context.Context, user user.User) (bool, error) {
	return
}
