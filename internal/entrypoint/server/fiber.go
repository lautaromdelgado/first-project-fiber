package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lautaromdelgado/first-project-fiber/internal/routes"
	"github.com/lautaromdelgado/first-project-fiber/internal/security/middlewares"
)

func InitFiber() *fiber.App {
	app := fiber.New()

	app.Use(middlewares.LoggingMiddleware)

	routes.InitRoutes(app)

	log.Fatal(app.Listen(":3000"))

	return app
}
