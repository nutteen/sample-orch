package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutteen/sample-orch/pkg/api/v1/hellohandler"
)

func RegisterRoutes(app *fiber.App) {
	helloHandler := hellohandler.NewHandler()

	routes := app.Group("/hello")
	routes.Get("/", helloHandler.HelloWorld)
}