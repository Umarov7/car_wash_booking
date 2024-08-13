package mongodb

import (
	"booking-service/models"
	"booking-service/storage"
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Notification() storage.INotificationStorage {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil
	}

	return NewNotificationRepo(client.Database("test"))
}

func TestCreateNotification(t *testing.T) {
	n := Notification()

	_, err := n.Create(context.Background(), &models.NewNotification{
		UserID:    "test",
		Title:     "test",
		Message:   "test",
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetNotification(t *testing.T) {
	n := Notification()

	_, err := n.Get(context.Background(), "66bb4faa2c8e3234ad957723") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}
