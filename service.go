package main

import (
	"github.com/google/uuid"
)

type Service struct {
	Repository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo}
}

func (s *Service) AddUser(user User) (uuid.UUID, error) {

	user.ID = uuid.New()

	id, err := s.Repository.CreateUser(user)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *Service) FindUser(id uuid.UUID) (*User, error) {

	user, err := s.Repository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
