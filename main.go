package main

import (
	"log"

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

	log.Fatal(app.Listen(":3000"))
}

func loggingMiddleware(c *fiber.Ctx) error {
	log.Printf("Solicitud recibida: %s %s\n", c.Method(), c.Path())
	return c.Next()
}
