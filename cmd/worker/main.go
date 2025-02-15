package main

import (
	"context"
	"insider-messaging/configs"
	"insider-messaging/internal/cache/redis"
	"insider-messaging/internal/database/mongodb"
	"insider-messaging/internal/message"
	"insider-messaging/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	appConfig, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load app config: %v", err)
	}

	mongodb.Connect(appConfig.Database.MongoURI)
	redis.Connect(appConfig.Cache.RedisAddr)
	messageService := message.NewService(message.NewRepository())
	messageWorker := worker.NewMessageWorker(messageService)

	ctx, cancel := context.WithCancel(context.Background())
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go messageWorker.Start(ctx)
	go messageWorker.HandleCommand(ctx)

	sig := <-sigChan
	log.Printf("Received signal: %v, initiating graceful shutdown...", sig)
	cancel()

	time.Sleep(time.Second * 2)
	log.Println("Gracefully shutdown.")
}
