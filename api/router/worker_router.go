package router

import (
	"insider-messaging/api/handler"

	"github.com/gofiber/fiber/v2"
)

func WorkerRouter(app fiber.Router) {
	app.Post("/worker/start", handler.StartWorker)
	app.Post("/worker/stop", handler.StopWorker)
}
