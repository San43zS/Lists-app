package notification

import (
	"Lists-app/internal/broker/rabbitMQ/api/notification"
	notification2 "Lists-app/internal/model/notification"
	notification3 "Lists-app/internal/service/api/notification"
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

func (s service) GetByUserId(ctx context.Context, Id int) (notification2.Notification, error) {
	return s.storage.GetByUserId(ctx, Id)
}

func (s service) GetList(ctx context.Context) ([]notification2.Notification, error) {
	return s.storage.GetList(ctx)
}

func (s service) Add(ctx context.Context, notification notification2.Notification) error {
	return s.storage.Add(ctx, notification)
}

func (s service) GetListWithTTL(ctx context.Context) ([]notification2.Notification, error) {
	return s.storage.GetListWithTTL(ctx)
}

func (s service) Delete(ctx context.Context, notification notification2.Notification) error {
	return s.storage.Delete(ctx, notification)
}
