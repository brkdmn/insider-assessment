package worker

import (
	"context"
	"insider-messaging/internal/cache/redis"
	"insider-messaging/internal/message"
	"log"
	"time"
)

type MessageWorker interface {
	Start(ctx context.Context)
	HandleCommand(ctx context.Context)
}

type messageWorker struct {
	messageService message.Service
	controlChan    chan string
	running        bool
}

func NewMessageWorker(messageService message.Service) MessageWorker {
	return &messageWorker{
		messageService: messageService,
		controlChan:    make(chan string),
		running:        true,
	}
}

func (w *messageWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Minute)
	defer ticker.Stop()
	running := true

	for {
		select {
		case <-ctx.Done():
			log.Println("Worker shutting down...")
			return
		case cmd := <-w.controlChan:
			switch cmd {
			case "start":
				running = true
				log.Println("Worker started")
			case "stop":
				running = false
				log.Println("Worker stopped")
			}
		case <-ticker.C:
			if running {
				log.Println("Sending messages...")
				w.messageService.SendPendingMessages(ctx)
			}
		}
	}
}

func (w *messageWorker) HandleCommand(ctx context.Context) {
	pubsub := redis.RedisClient.Subscribe(context.Background(), "worker_control")
	defer pubsub.Close()

	for {
		select {
		case <-ctx.Done():
			log.Println("Command listener shutting down...")
			return
		default:
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Println("Redis message not received:", err)
				continue
			}

			if msg.Payload == "start" || msg.Payload == "stop" {
				w.controlChan <- msg.Payload
			}
		}
	}
}
