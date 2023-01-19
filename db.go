package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewReposiory(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateUser(user User) (uuid.UUID, error) {
	res := r.DB.Create(&user)
	if res.Error != nil {
		return uuid.Nil, res.Error
	}

	return user.ID, nil
}

func (r *Repository) GetUser(id uuid.UUID) (*User, error) {
	var user User
	if err := r.DB.First(&user, id).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, err
	}
	return &user, nil
}
