package service

import (
	"github.com/google/uuid"
	"github.com/naopin/coin-be/graph/model"
	"github.com/naopin/coin-be/graph/repository"
)

type IUserService interface {
	CreateItem(input model.CreateUserInput) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type UserService struct {
	repository repository.IUserRepository // インターフェース型に変更
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository}
}

func (us *UserService) CreateItem(input model.CreateUserInput) (*model.User, error) {
	uuid, _ := uuid.NewUUID()
	user := model.User{
		ID:       uuid.String(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(input.Password),
	}
	return us.repository.Create(user)
}

func (us *UserService) GetUserByEmail(email string) (*model.User, error) {

	return us.repository.GetUserByEmail(email)
}
