package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lautaromdelgado/first-project-fiber/internal/User/handler"
	usecase "github.com/lautaromdelgado/first-project-fiber/internal/User/use_case"
)

func InitRoutes(app *fiber.App) {
	userUseCases := usecase.NewUserUseCase()
	userHandler := handler.NewUserHandler(userUseCases)

	app.Post("/usuario/crear", userHandler.Create)
}
