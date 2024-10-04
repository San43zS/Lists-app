package service

import (
	"Lists-app/internal/storage"
)

const (
	uniqueViolationErr = "23505"
)

type Service struct {
	storage *storage.Storage
}

func New(repos storage.Storage) *Service {

	return &Service{}
}
