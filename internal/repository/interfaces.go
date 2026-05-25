package repository

import "booking-system/lvl1/internal/model"

type UserRepo interface {
	Create(username, password string, role model.UserType) error
	GetByUsername(username string) (*model.User, error)
	GetByID(userID string) (*model.User, error)
}
