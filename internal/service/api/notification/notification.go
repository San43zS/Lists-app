package notification

import (
	"Lists-app/internal/model/notification"
	"context"
)

type Notification interface {
	GetNotificationByUserId(ctx context.Context, Id int) (notification.Notification, error)
	GetAllNotificationsTTLoff(ctx context.Context) ([]notification.Notification, error)
	AddNotification(ctx context.Context, notification notification.Notification) error
	GetAllNotificationsTTLon(ctx context.Context) ([]notification.Notification, error)
	DeleteNotification(ctx context.Context, notification notification.Notification) error
}
