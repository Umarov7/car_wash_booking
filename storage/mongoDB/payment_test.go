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

func Payment() storage.IPaymentStorage {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil
	}

	return NewPaymentRepo(client.Database("test"))
}

func TestCreatePayment(t *testing.T) {
	p := Payment()

	_, err := p.Create(context.Background(), &models.NewPayment{
		BookingId:     "test",
		Amount:        100,
		Status:        "test",
		PaymentMethod: "test",
		TransactionId: "test",
		CreatedAt:     time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetPayment(t *testing.T) {
	p := Payment()

	_, err := p.Get(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestFetchPayment(t *testing.T) {
	p := Payment()

	_, err := p.Fetch(context.Background(), 1, 10)
	if err != nil {
		t.Error(err)
	}
}
