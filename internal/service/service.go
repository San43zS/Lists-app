package service

import (
	"Lists-app/internal/service/api/notification"
	"Lists-app/internal/service/api/user"
	"Lists-app/internal/storage"
)

const (
	uniqueViolationErr = "23505"
)

type Service interface {
	User() user.User
	Notification() notification.Notification
}

type service struct {
	storage      *storage.Storage
	user         user.User
	notification notification.Notification
}

func New(repos storage.Storage) *Service {

	return &service{
		storage:      repos,
		user:         user.New(repos),
		notification: notification.New(repos),
	}
}
