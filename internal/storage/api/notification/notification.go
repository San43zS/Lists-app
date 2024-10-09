package notification

import (
	"Lists-app/internal/model/notification"
	"context"
)

type Notification interface {
	GetByUserId(ctx context.Context, Id int) (notification.Notification, error)
	GetList(ctx context.Context) ([]notification.Notification, error)
	Add(ctx context.Context, notification notification.Notification) error
	GetListWithTTL(ctx context.Context) ([]notification.Notification, error)
	Delete(ctx context.Context, notification notification.Notification) error
}
