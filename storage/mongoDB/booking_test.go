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

func Booking() storage.IBookingStorage {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil
	}

	return NewBookingRepo(client.Database("test"))
}

func TestCreateBooking(t *testing.T) {
	b := Booking()

	_, err := b.Create(context.Background(), &models.NewBooking{
		UserId:      "test",
		ProviderId:  "test",
		ServiceId:   "test",
		Status:      "test",
		ScheduledAt: time.Now().Format(time.RFC3339),
		Location: models.Location{
			Address:   "test",
			City:      "test",
			Country:   "test",
			Latitude:  1.0,
			Longitude: 1.0,
		},
		TotalPrice: 1.0,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetBooking(t *testing.T) {
	b := Booking()

	_, err := b.Get(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateBooking(t *testing.T) {
	b := Booking()

	err := b.Update(context.Background(), &models.NewBookingData{
		Id:          "", // ObjectIdHex
		Status:      "pending",
		ScheduledAt: time.Now().Format(time.RFC3339),
		Location: models.Location{
			Address:   "test",
			City:      "test",
			Country:   "test",
			Latitude:  1.0,
			Longitude: 1.0,
		},
		TotalPrice: 1.0,
		UpdatedAt:  time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestCancelBooking(t *testing.T) {
	b := Booking()

	err := b.Cancel(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestFetchBooking(t *testing.T) {
	b := Booking()

	_, err := b.Fetch(context.Background(), 1, 10)
	if err != nil {
		t.Error(err)
	}
}
