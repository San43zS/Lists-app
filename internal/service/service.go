package service

import (
	"Lists-app/internal/model/user"
	"Lists-app/internal/storage"
	"github.com/lib/pq"
)

const (
	uniqueViolationErr = "23505"
)

type Authorization interface {
	Registration(user user.User) error
	Authenticate(user user.User) error
	GenerateToken(user user.User) (string, error)
}

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
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == uniqueViolationErr {
			return nil
		}
	}
	return nil
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	storage *storage.Storage
	Authorization
}

func New(repos *storage.Storage) *Service {
	return &Service{}
}

func (s *Service) VerifyUser(user user.User) error {