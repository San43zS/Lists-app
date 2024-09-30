package service

import (
	"Lists-app/internal/model/user"
	"Lists-app/internal/service/auth"
	"Lists-app/internal/storage"
	"github.com/lib/pq"
)

const (
	uniqueViolationErr = "23505"
)

func (s *Service) Authenticate(user user.User) error {
	err := s.storage.VerifyUser(user)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == uniqueViolationErr {
			return nil
		}
	}
	return err
}

func (s *Service) Registration(user user.User) error {
	err := s.storage.InsertUser(user)
	if err != nil {
		return err
	}
	return nil
}

type Service struct {
	storage *storage.Storage
	auth.Autorization
}

func New(repos storage.Storage) *Service {
	return &Service{}
}
