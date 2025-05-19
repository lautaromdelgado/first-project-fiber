package handler

import (
	"github.com/gofiber/fiber/v2"
	useCase "github.com/lautaromdelgado/first-project-fiber/internal/User/use_case"
)

type IUserHandler interface {
	Create(*fiber.Ctx) error
}

type UserHandler struct {
	useCase useCase.IUserUseCase
}

func NewUserHandler(useCase useCase.IUserUseCase) IUserHandler {
	return &UserHandler{
		useCase: useCase,
	}
}
