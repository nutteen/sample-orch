package main

import (
	"gitdev.devops.krungthai.com/open-api/backend/common/core-lib/core/logger"
	middlewarelogger "gitdev.devops.krungthai.com/open-api/backend/common/core-lib/core/middlewares/logger"
	"gitdev.devops.krungthai.com/open-api/backend/common/core-lib/core/middlewares/usercontext"
	"gitdev.devops.krungthai.com/open-api/backend/common/core-lib/core/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
