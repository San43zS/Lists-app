package notification

import (
	notification2 "Lists-app/internal/model/notification"
	notification3 "Lists-app/internal/service/api/notification"
	"Lists-app/internal/storage/api/notification"
	"context"
)

type service struct {
	storage notification.Notification
}

func New(storage notification3.Notification) notification3.Notification {
	return &service{
		storage: storage,
	}
}

func (s service) GetNotificationByUserId(ctx context.Context, Id int) (notification2.Notification, error) {
	return s.storage.GetNotificationByUserId(ctx, Id)
}

func (s service) GetAllNotificationsTTLoff(ctx context.Context) ([]notification2.Notification, error) {
	return s.storage.GetAllNotificationsTTLoff(ctx)
}

func (s service) AddNotification(ctx context.Context, notification notification2.Notification) error {
	return s.storage.AddNotification(ctx, notification)
}

func (s service) GetAllNotificationsTTLon(ctx context.Context) ([]notification2.Notification, error) {
	return s.storage.GetAllNotificationsTTLon(ctx)
}

func (s service) DeleteNotification(ctx context.Context, notification notification2.Notification) error {
	return s.storage.DeleteNotification(ctx, notification)
}
