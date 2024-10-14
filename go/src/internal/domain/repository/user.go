package repository

import (
	"github.com/kynmh69/mormorare/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) GetUsers() ([]domain.User, error) {
	var users []domain.User
	if err := u.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil

}

func (u *UserRepository) CreateUser(user *domain.User) error {
	if err := u.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) UpdateUser(user *domain.User) error {
	if err := u.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	var user domain.User
	if err := u.db.Where("username = ?", username).Take(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) DeleteUser(user *domain.User) error {
	if err := u.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}
