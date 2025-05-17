package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func InitFiber() *fiber.App {
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))

	return app
}
