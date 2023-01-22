package main

import (
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewReposiory(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user User) (uuid.UUID, error) {
	res := r.db.Create(&user)
	if res.Error != nil {
		return uuid.Nil, res.Error
	}

	return user.ID, nil
}

func (r *UserRepository) GetUser(id uuid.UUID) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			return nil, fmt.Errorf("user with id %s not found", id)
		}
		return nil, err
	}
	return &user, nil
}
