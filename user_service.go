package main

import (
	"github.com/google/uuid"
)

type UserService struct {
	repository *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) AddUser(user User) (uuid.UUID, error) {

	user.ID = uuid.New()

	id, err := s.repository.CreateUser(user)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *UserService) FindUser(id uuid.UUID) (*User, error) {

	user, err := s.repository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
