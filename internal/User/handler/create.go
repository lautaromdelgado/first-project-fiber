package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lautaromdelgado/first-project-fiber/internal/User/dto"
)

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var user dto.UserRequest

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parseando el cuerpo de la solicitud",
		})
	}

	userResponse, err := h.useCase.CreateUser(&user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "ooops, algo salió mal",
		})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"message": "Usuario creado con éxito",
		"user":    userResponse,
	})
}
