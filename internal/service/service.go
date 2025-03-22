package service

import (
	"notify-service/internal/broker"
	"notify-service/internal/service/api/notification"
	"notify-service/internal/service/api/user"
	notification2 "notify-service/internal/service/notification"
	user2 "notify-service/internal/service/user"
	"notify-service/internal/storage"
)

type Service interface {
	User() user.User
	Notification() notification.Notification
}

type service struct {
	storage      storage.Storage
	user         user.User
	notification notification.Notification
}

func New(repos storage.Storage, broker broker.Broker) Service {
	return &service{
		storage:      repos,
		user:         user2.New(repos.User()),
		notification: notification2.New(broker),
	}
}

func (s *service) User() user.User {
	return s.user
}

func (s *service) Notification() notification.Notification {
	return s.notification
}
