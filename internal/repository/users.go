package repository

import (
	"booking-system/lvl1/internal/config"
	"booking-system/lvl1/internal/model"
)

type UserRepository struct {
	db *config.DB
}

func NewUserRepository(db *config.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(username, password string, role model.UserType) error {
	user := model.User{
		Username: username,
		Password: password,
		Role:     role,
	}
	return r.db.Create(&user).Error
}

func (r *UserRepository) GetByUsername(username string) (*model.User, error) {
	var user model.User
	result := r.db.
		Preload("PersonalDetails").
		Where("username = ?", username).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetByID(userID string) (*model.User, error) {
	var user model.User
	result := r.db.
		Select("id", "username", "password", "role").
		Where("id = ?", userID).
		First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
