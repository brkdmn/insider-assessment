package main

import (
	"insider-messaging/api/router"
	"insider-messaging/configs"
	"insider-messaging/internal/cache/redis"
	"insider-messaging/internal/database/mongodb"
	"insider-messaging/internal/message"
	"log"

	_ "insider-messaging/docs"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

// @title Insider Messaging API
// @version 2.0
// @description Automatic Sending Message Service
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {

	appConfig, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load app config: %v", err)
	}

	mongodb.Connect(appConfig.Database.MongoURI)
	redis.Connect(appConfig.Cache.RedisAddr)

	app := fiber.New()
	api := app.Group("/api")

	router.MessageRouter(api, message.NewService(message.NewRepository()))
	router.WorkerRouter(api)
	app.Get("/swagger/*", swagger.HandlerDefault)

	log.Println("API started: http://localhost:8080")
	app.Listen(":8080")
}
