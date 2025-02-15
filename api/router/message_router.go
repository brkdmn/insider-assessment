package router

import (
	"insider-messaging/api/handler"
	"insider-messaging/internal/message"

	"github.com/gofiber/fiber/v2"
)

func MessageRouter(app fiber.Router, service message.Service) {
	app.Get("/messages", func(c *fiber.Ctx) error {
		return handler.GetMessages(c, service)
	})
}
