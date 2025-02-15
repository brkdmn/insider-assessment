package router

import (
	"insider-messaging/api/handler"

	"github.com/gofiber/fiber/v2"
)

func WorkerRouter(app fiber.Router) {
	app.Post("/start", handler.StartWorker)
	app.Post("/stop", handler.StopWorker)
}
