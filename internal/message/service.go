package message

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"insider-messaging/internal/cache/redis"
	"log"
	"net/http"
	"time"
)

const (
	MaxMessagesToSend = 2
	MaxMessageLength  = 160
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

		if len(msg.Content) > MaxMessageLength {
			log.Printf("Message content is too long: %s", msg.Content)
			continue
		}

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
		if err := cacheMessage(msg.ID, response.MessageID); err != nil {
			log.Printf("Failed to cache message: %v", err)
		}
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
		data, err := redis.RedisClient.Get(ctx, key).Bytes()
		if err != nil {
			log.Printf("Redis read error: %v", err)
			continue
		}

		var sendingMessage SendingMessage
		if err := json.Unmarshal(data, &sendingMessage); err != nil {
			log.Printf("JSON unmarshal error: %v", err)
			continue
		}
		sendingMessages = append(sendingMessages, sendingMessage)
	}

	return sendingMessages, nil
}

func sendMessageToWebhook(phone, content string) (*MessageResponse, error) {
	const webhookURL = "https://webhook.site/a70ce157-8ab2-4fc4-af79-2a9eaa8e74d2"

	payload := struct {
		To      string `json:"to"`
		Content string `json:"content"`
	}{
		To:      phone,
		Content: content,
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal request body failed: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("webhook request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response MessageResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	return &response, nil
}

func cacheMessage(id string, messageId string) error {
	const (
		keyPrefix     = "messages:id:"
		cacheDuration = 24 * time.Hour
	)

	msgData := SendingMessage{
		MessageID: messageId,
		SentAt:    time.Now(),
	}

	jsonData, err := json.Marshal(msgData)
	if err != nil {
		return fmt.Errorf("marshal message data failed: %w", err)
	}

	ctx := context.Background()
	key := keyPrefix + id
	if err := redis.RedisClient.Set(ctx, key, jsonData, cacheDuration).Err(); err != nil {
		return fmt.Errorf("redis set failed: %w", err)
	}

	return nil
}
