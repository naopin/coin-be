package repository

import (
	"github.com/naopin/coin-be/graph/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(user model.User) (*model.User, error) {

	err := ur.db.Create(user).Error
	return nil, err
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	var user *model.User
	err := ur.db.First(&user, "email = ?", email).Find(&user).Error

	return user, err

}
