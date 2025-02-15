package handler

import (
	"context"
	"insider-messaging/internal/cache/redis"
	"log"

	"github.com/gofiber/fiber/v2"
)

// StartWorker godoc
// @Summary Start worker
// @Description Start worker with redis channel
// @Tags Worker
// @Router /api/worker/start [post]
func StartWorker(c *fiber.Ctx) error {
	err := redis.RedisClient.Publish(context.Background(), "worker_control", "start").Err()
	if err != nil {
		log.Println("Worker start message not sent:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Worker start message not sent"})
	}

	return c.JSON(fiber.Map{"message": "Worker started."})
}

// StopWorker godoc
// @Summary Stop worker
// @Description Stop worker with redis channel
// @Tags Worker
// @Router /api/worker/stop [post]
func StopWorker(c *fiber.Ctx) error {
	err := redis.RedisClient.Publish(context.Background(), "worker_control", "stop").Err()
	if err != nil {
		log.Println("Worker stop message not sent:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Worker stop message not sent"})
	}

	return c.JSON(fiber.Map{"message": "Worker stopped."})
}
