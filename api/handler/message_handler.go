package handler

import (
	"context"
	"insider-messaging/internal/message"

	"github.com/gofiber/fiber/v2"
)

// GetMessages godoc
// @Summary Sending Messages
// @Description Get sending messages
// @Tags Messages
// @Accept  json
// @Produce json
// @Success 200 {array} message.SendingMessage
// @Router /api/messages [get]
func GetMessages(c *fiber.Ctx, service message.Service) error {
	messages, err := service.GetSendingMessages(context.Background())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Mesajlar getirilemedi"})
	}
	return c.JSON(messages)
}
