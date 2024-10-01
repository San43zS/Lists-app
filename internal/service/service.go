package service

import (
	"Lists-app/internal/service/auth"
	"Lists-app/internal/service/notification"
	"Lists-app/internal/storage"
)

const (
	uniqueViolationErr = "23505"
)

type Service struct {
	storage *storage.Storage
	auth.Autorization
	notification.Notification
}

func New(repos storage.Storage) *Service {
	return &Service{}
}
