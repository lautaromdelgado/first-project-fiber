package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautaromdelgado/first-project-fiber/internal/User/dto"
)

func InitRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"stattus": "success",
			"message": "Hello, World!",
		})
	})

	app.Post("/api/usuarios", func(c *fiber.Ctx) error {
		user := new(dto.UserRequest)

		if err := c.BodyParser(user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
				"status":  "error",
				"message": "oooops, algo salió mal",
			})
		}

		userResponse := dto.UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		}

		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "success",
			"message": "usuario creado",
			"user":    userResponse,
		})
	})
}
