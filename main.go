package main

import (
	"github.com/lautaromdelgado/first-project-fiber/internal/entrypoint/server"
	"github.com/lautaromdelgado/first-project-fiber/internal/security/middlewares"
)

func main() {
	fiber := server.InitFiber()

	fiber.Use(middlewares.LoggingMiddleware)
}
