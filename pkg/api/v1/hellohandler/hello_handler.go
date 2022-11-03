package hellohandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutteen/sample-orch/pkg/model"
)

type HelloHandler struct {

}

func NewHandler() *HelloHandler {
	return &HelloHandler{}
}

func (h HelloHandler) HelloWorld(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&model.HelloResponse{
		GreetingText: "Hello Friend",
	})
}