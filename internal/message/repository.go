package message

import (
	"context"
	"log"
	"time"

	"insider-messaging/internal/database/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GetPendingMessages(ctx context.Context, limit int) ([]Message, error)
	AddMessage(ctx context.Context, msg Message) error
	MarkMessageAsSent(ctx context.Context, id string) error
}

type repository struct {
	collection *mongo.Collection
}

func NewRepository() Repository {
	return &repository{
		collection: mongodb.MongoDB.Collection("messages"),
	}
}

func (r *repository) GetPendingMessages(ctx context.Context, limit int) ([]Message, error) {
	filter := bson.M{"is_sent": false}
	opts := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}}).SetLimit(int64(limit))
	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}

	var messages []Message
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func (r *repository) AddMessage(ctx context.Context, msg Message) error {
	msg.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, msg)
	return err
}

func (r *repository) MarkMessageAsSent(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid ObjectId")
		return err
	}
	_, err = r.collection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": bson.M{"is_sent": true, "sent_at": time.Now()}})
	return err
}
