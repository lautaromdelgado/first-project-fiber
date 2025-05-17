package middlewares

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(c *fiber.Ctx) error {
	log.Printf("Solicitud recibida: %s %s\n", c.Method(), c.Path())
	return c.Next()
}
