package tests

import (
	"booking-system/lvl1/internal/model"
	"booking-system/lvl1/internal/repository"
	"booking-system/lvl1/internal/service"
	"errors"
	"testing"
)

type mockUserRepo struct {
	user      *model.User
	getErr    error
	createErr error
}

func (m *mockUserRepo) Create(username, password string, role model.UserType) error {
	return m.createErr
}

func (m *mockUserRepo) GetByUsername(username string) (*model.User, error) {
	return m.user, m.getErr
}

func (m *mockUserRepo) GetByID(userID string) (*model.User, error) {
	return m.user, m.getErr
}

var _ repository.UserRepo = (*mockUserRepo)(nil)

var errNotFound = errors.New("not found")

func TestCreateUser_DefaultsRoleToUser(t *testing.T) {
	repo := &mockUserRepo{getErr: errNotFound}
	svc := service.NewUserService(repo)

	req := model.CreateUserRequest{Username: "mark", Password: "password123"}
	svc.CreateUser(req)

	if req.Role != "" {
		t.Errorf("original request should not be mutated")
	}
}

func TestCreateUser_ReturnsErrUserExists(t *testing.T) {
	repo := &mockUserRepo{user: &model.User{Username: "mark"}}
	svc := service.NewUserService(repo)

	req := model.CreateUserRequest{Username: "mark", Password: "password123"}
	err := svc.CreateUser(req)

	if err != service.ErrUserExists {
		t.Errorf("expected ErrUserExists, got %v", err)
	}
}

func TestCreateUser_Success(t *testing.T) {
	repo := &mockUserRepo{getErr: errNotFound, createErr: nil}
	svc := service.NewUserService(repo)

	req := model.CreateUserRequest{Username: "mark", Password: "password123"}
	err := svc.CreateUser(req)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
