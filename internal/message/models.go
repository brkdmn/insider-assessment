package message

import "time"

type Message struct {
	ID          string     `bson:"_id"`
	PhoneNumber string     `bson:"phone_number"`
	Content     string     `bson:"content"`
	IsSent      bool       `bson:"is_sent"`
	CreatedAt   time.Time  `bson:"created_at"`
	SentAt      *time.Time `bson:"sent_at"`
}

type MessageResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

type SendingMessage struct {
	MessageID string    `json:"messageId"`
	SentAt    time.Time `bson:"sent_at"`
}
