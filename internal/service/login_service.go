package service

import "github.com/stephenZ22/mini_dash/internal/repository"

type LoginService interface {
	LoginByPassword(username, password string) (bool, error)
}

type loginService struct {
	repo repository.UserRepository
}

func NewLoginService(repo repository.UserRepository) LoginService {
	return &loginService{
		repo: repo,
	}
}
func (lg *loginService) LoginByPassword(username, password string) (bool, error) {
	ok, err := lg.repo.CheckPassword(username, password)
	return ok, err
}
