package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	app := fiber.New()

	app.Use(loggingMiddleware)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"stattus": "success",
			"message": "Hello, World!",
		})
	})

	app.Post("/api/usuarios", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(map[string]interface{}{
				"status":  "error",
				"message": "oooops, algo salió mal",
			})
		}
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"status":  "success",
			"message": "usuario creado",
			"user":    user,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

func loggingMiddleware(c *fiber.Ctx) error {
	log.Printf("Solicitud recibida: %s %s\n", c.Method(), c.Path())
	return c.Next()
}
