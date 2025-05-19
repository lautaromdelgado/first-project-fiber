package usecase

import (
	"errors"

	"github.com/lautaromdelgado/first-project-fiber/internal/User/dto"
)

func (userUC *UserUseCase) CreateUser(user *dto.UserRequest) (*dto.UserResponse, error) {
	if user.ID <= 0 {
		return nil, errors.New("ID is required")
	}
	if user.Username == "" {
		return nil, errors.New("Username is required")
	}
	if user.Email == "" {
		return nil, errors.New("Email is required")
	}
	userResponse := dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return &userResponse, nil
}
