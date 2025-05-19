package usecase

import "github.com/lautaromdelgado/first-project-fiber/internal/User/dto"

type IUserUseCase interface {
	CreateUser(user *dto.UserRequest) (*dto.UserResponse, error)
}

type UserUseCase struct{}

func NewUserUseCase() IUserUseCase {
	return &UserUseCase{}
}
