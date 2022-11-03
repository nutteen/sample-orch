package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/nutteen/png-core/core/logger"
	middlewarelogger "github.com/nutteen/png-core/core/middlewares/logger"
	"github.com/nutteen/png-core/core/middlewares/usercontext"
	"github.com/nutteen/png-core/core/validator"
	"github.com/nutteen/sample-orch/pkg/router"
)


func main() {

	logger.InitializeLogger(logger.LoggerConfig{
		IsProductionMode: true,
	})

	validator.InitializeValidator()

	app := fiber.New()

	// Setup middlewares
	app.Use(requestid.New())
	app.Use(usercontext.New())
	app.Use(middlewarelogger.NewLogger(logger.Log, middlewarelogger.ConfigDefault))
	// Registers routes
	router.RegisterRoutes(app)

	app.Listen(":3000")
}