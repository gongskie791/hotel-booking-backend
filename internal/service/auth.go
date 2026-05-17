package service

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/repository"
	"booking-system/lvl1/internal/util"
	"errors"
)

var ErrInvalidCredentials = errors.New("invalid username or password")
var ErrInvalidToken = errors.New("invalid token")
var ErrUnauthorized = errors.New("unauthorized")

type AuthService struct {
	userRepository *repository.UserRepository
}

func NewAuthService(userRepository *repository.UserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (s *AuthService) Login(username, password string) (*model.LoginResponse, string, error) {
	user, err := s.userRepository.GetByUsername(username)
	if err != nil {
		return nil, "", ErrInvalidCredentials
	}

	if !util.CheckPasswordHash(password, user.Password) {
		return nil, "", ErrInvalidCredentials
	}

	accessToken, err := util.GenerateToken(user.ID, user.Role, util.AccessToken)
	if err != nil {
		return nil, "", err
	}

	refreshToken, err := util.GenerateToken(user.ID, user.Role, util.RefreshToken)
	if err != nil {
		return nil, "", err
	}

	userInfo := model.LoginUserInfo{ID: user.ID}
	if user.PersonalDetails != nil {
		userInfo.Name = user.PersonalDetails.Name
		userInfo.Email = user.PersonalDetails.Email
	}

	return &model.LoginResponse{
		Token: accessToken,
		User:  userInfo,
	}, refreshToken, nil
}

func (s *AuthService) RefreshToken(incomingRefreshToken string) (accessToken string, refreshToken string, err error) {
	claims, err := util.ValidateToken(incomingRefreshToken)
	if err != nil {
		return "", "", ErrInvalidToken
	}

	if claims.TokenType != util.RefreshToken {
		return "", "", ErrInvalidToken
	}

	user, err := s.userRepository.GetByID(claims.UserID)
	if err != nil {
		return "", "", ErrUnauthorized
	}

	newAccessToken, err := util.GenerateToken(user.ID, user.Role, util.AccessToken)
	if err != nil {
		return "", "", err
	}

	newRefreshToken, err := util.GenerateToken(user.ID, user.Role, util.RefreshToken)
	if err != nil {
		return "", "", err
	}

	return newAccessToken, newRefreshToken, nil
}
