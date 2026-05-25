package tests

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/service"
	"booking-system/lvl1/internal/util"
	"testing"
)

func TestLogin_UserNotFound(t *testing.T) {
	repo := &mockUserRepo{user: nil, getErr: errNotFound}
	svc := service.NewAuthService(repo)

	_, _, err := svc.Login("unknown", "password")

	if err != service.ErrInvalidCredentials {
		t.Errorf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestLogin_WrongPassword(t *testing.T) {
	hashed, _ := util.HashPassword("correctpassword")
	repo := &mockUserRepo{
		user: &model.User{ID: "1", Username: "mark", Password: hashed, Role: model.UserTypeUser},
	}
	svc := service.NewAuthService(repo)

	_, _, err := svc.Login("mark", "wrongpassword")

	if err != service.ErrInvalidCredentials {
		t.Errorf("expected ErrInvalidCredentials, got %v", err)
	}
}

func TestLogin_Success(t *testing.T) {
	hashed, _ := util.HashPassword("correctpassword")
	repo := &mockUserRepo{
		user: &model.User{ID: "1", Username: "mark", Password: hashed, Role: model.UserTypeUser},
	}
	svc := service.NewAuthService(repo)

	resp, refreshToken, err := svc.Login("mark", "correctpassword")

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if resp.Token == "" {
		t.Error("expected access token, got empty string")
	}
	if refreshToken == "" {
		t.Error("expected refresh token, got empty string")
	}
}
