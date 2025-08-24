package repository

import (
	"errors"

	"github.com/stephenZ22/mini_dash/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	GetUserByID(id uint) (*model.User, error)
	UpdateUser(user *model.User) error
	DeleteUser(id uint) error
	CheckPassword(username, password string) (ok *model.User, err error)
	List(page_size, page_num int) ([]model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByID(id uint) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) CheckPassword(username, password string) (*model.User, error) {
	var user model.User

	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	if user.Password == password {
		return &user, nil
	}

	return nil, errors.New("invalid password")
}

func (r *userRepository) List(page_size, page_num int) ([]model.User, error) {
	var users []model.User

	if page_size < 0 || page_num < 0 {
		return nil, gorm.ErrInvalidValue
	}

	if page_num == 0 {
		page_num = 1
	}

	if page_size == 0 {
		page_size = 20
	}
	if err := r.db.Offset((page_num - 1) * page_size).Limit(page_size).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
