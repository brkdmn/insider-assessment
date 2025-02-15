package message

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"insider-messaging/internal/cache/redis"
	"log"
	"net/http"
	"time"
)

const (
	MaxMessagesToSend = 2
)

type Service interface {
	SendPendingMessages(ctx context.Context)
	GetSendingMessages(ctx context.Context) ([]SendingMessage, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) SendPendingMessages(ctx context.Context) {
	messages, err := s.repo.GetPendingMessages(ctx, MaxMessagesToSend)
	if err != nil {
		log.Println("Error: Failed to get pending messages", err)
		return
	}

	for _, msg := range messages {
		lockKey := "lock:" + msg.ID
		acquired := redis.AcquireLock(ctx, lockKey, 2*time.Second)
		if !acquired {
			log.Printf("Message is already being processed by another process: %s", msg.ID)
			continue
		}

		response, err := sendMessageToWebhook(msg.PhoneNumber, msg.Content)
		if err != nil {
			log.Printf("Error: Failed to send message to webhook: %s", err)
			continue
		}

		s.repo.MarkMessageAsSent(ctx, msg.ID)
		cacheMessage(msg.ID, response.MessageID)
		redis.ReleaseLock(ctx, lockKey)
		log.Printf("Message sent: %s, ID: %s", msg.Content, response.MessageID)
	}
}

func (s *service) GetSendingMessages(ctx context.Context) ([]SendingMessage, error) {
	keys, err := redis.RedisClient.Keys(ctx, "messages:id:*").Result()
	if err != nil {
		return nil, err
	}

	var sendingMessages []SendingMessage
	for _, key := range keys {
		var sendingMessage SendingMessage
		err := redis.RedisClient.Get(ctx, key).Scan(&sendingMessage)
		if err != nil {
			continue
		}
		sendingMessages = append(sendingMessages, sendingMessage)
	}

	return sendingMessages, nil
}

func sendMessageToWebhook(phone, content string) (*MessageResponse, error) {
	url := "https://webhook.site/18623911-b0f8-42ee-adf8-f2d4175d57aa"
	requestBody, _ := json.Marshal(map[string]string{
		"to":      phone,
		"content": content,
	})

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Webhook request failed:", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusAccepted {
		var response MessageResponse
		json.NewDecoder(resp.Body).Decode(&response)
		return &response, nil
	}
	return nil, errors.New("failed to send message to webhook")
}

func cacheMessage(id string, messageId string) {
	ctx := context.Background()
	msgData := SendingMessage{
		MessageID: messageId,
		SentAt:    time.Now(),
	}

	jsonData, _ := json.Marshal(msgData)
	key := "messages:id:" + id
	err := redis.RedisClient.Set(ctx, key, jsonData, 24*time.Hour).Err()
	if err != nil {
		log.Println("Redis write error:", err)
	}
}
