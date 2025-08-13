package users

type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id uint) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id uint) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *userService) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(user *User) error {
	return s.repo.UpdateUser(user)
}

func (s *userService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}
