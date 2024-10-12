package service

import (
	"Lists-app/internal/service/api/notification"
	"Lists-app/internal/service/api/user"
	notification2 "Lists-app/internal/service/notification"
	user2 "Lists-app/internal/service/user"
	"Lists-app/internal/storage"
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

func New(repos storage.Storage) Service {
	return &service{
		storage:      repos,
		user:         user2.New(repos.User()),
		notification: notification2.New(repos.Notification()),
	}
}

func (s *service) User() user.User {
	return s.user
}

func (s *service) Notification() notification.Notification {
	return s.notification
}
