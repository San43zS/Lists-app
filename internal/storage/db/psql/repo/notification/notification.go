package notification

import (
	notification3 "Lists-app/internal/model/notification"
	notification2 "Lists-app/internal/storage/api/notification"
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) notification2.Notification {

	return repository{
		db: db,
	}
}

func (r repository) Add(ctx context.Context, notification notification3.Notification) error {
	return nil
}

func (r repository) GetByUserId(ctx context.Context, Id int) (notification3.Notification, error) {
	return notification3.Notification{}, nil
}

func (r repository) GetListWithTTL(ctx context.Context) ([]notification3.Notification, error) {
	return nil, nil
}

func (r repository) GetList(ctx context.Context) ([]notification3.Notification, error) {
	return nil, nil
}

func (r repository) Delete(ctx context.Context, notification notification3.Notification) error {
	return nil
}
