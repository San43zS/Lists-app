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

func (r repository) AddNotification(ctx context.Context, notification notification3.Notification) error {
	return nil
}

func (r repository) GetNotificationByUserId(ctx context.Context, Id int) (notification3.Notification, error) {
	return notification3.Notification{}, nil
}

func (r repository) GetAllNotificationsTTLon(ctx context.Context) ([]notification3.Notification, error) {
	return nil, nil
}

func (r repository) GetAllNotificationsTTLoff(ctx context.Context) ([]notification3.Notification, error) {
	return nil, nil
}

func (r repository) DeleteNotification(ctx context.Context, notification notification3.Notification) error {
	return nil
}
