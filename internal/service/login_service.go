package service

import (
	"github.com/stephenZ22/mini_dash/internal/model"
	"github.com/stephenZ22/mini_dash/internal/repository"
)

type LoginService interface {
	LoginByPassword(username, password string) (*model.User, error)
}

type loginService struct {
	repo repository.UserRepository
}

func NewLoginService(repo repository.UserRepository) LoginService {
	return &loginService{
		repo: repo,
	}
}
func (lg *loginService) LoginByPassword(username, password string) (*model.User, error) {
	user, err := lg.repo.CheckPassword(username, password)
	return user, err
}
