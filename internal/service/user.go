package service

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/repository"
	"booking-system/lvl1/internal/util"
	"errors"
)

var ErrUserExists = errors.New("user already exists")

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) CreateUser(req model.CreateUserRequest) error {
	if req.Role == "" {
		req.Role = model.UserTypeUser
	}

	user, err := s.userRepository.GetByUsername(req.Username)
	if err == nil && user != nil {
		return ErrUserExists
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	return s.userRepository.Create(req.Username, hashedPassword, req.Role)
}
