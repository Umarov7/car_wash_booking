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

func Review() storage.IReviewStorage {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil
	}

	return NewReviewRepo(client.Database("test"))
}

func TestCreateReview(t *testing.T) {
	r := Review()

	_, err := r.Create(context.Background(), &models.NewReview{
		BookingId:  "test",
		UserId:     "test",
		ProviderId: "test",
		Rating:     1,
		Comment:    "test",
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetReview(t *testing.T) {
	r := Review()

	_, err := r.Get(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateReview(t *testing.T) {
	r := Review()

	err := r.Update(context.Background(), &models.NewReviewData{
		Id:        "", // ObjectIdHex
		Rating:    5,
		Comment:   "test",
		UpdatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteReview(t *testing.T) {
	r := Review()

	err := r.Delete(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestFetchReview(t *testing.T) {
	r := Review()

	_, err := r.Fetch(context.Background(), 1, 10)
	if err != nil {
		t.Error(err)
	}
}
