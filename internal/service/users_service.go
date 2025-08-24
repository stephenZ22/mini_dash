package service

import (
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	ListUsers(pageSize, pageNum int) ([]model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *model.User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *userService) ListUsers(pageSize, pageNum int) ([]model.User, error) {
	if pageSize < 0 || pageNum < 0 {
		return nil, gorm.ErrInvalidValue
	}

	if pageNum == 0 {
		pageNum = 1
	}

	if pageSize == 0 {
		pageSize = 10 // Default page size
	}

	users, err := s.repo.List(pageSize, pageNum)
	if err != nil {
		return nil, err
	}
	return users, nil
}
