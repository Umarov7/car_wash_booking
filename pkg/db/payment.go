package db

import (
	"booking-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func paymentData() []*models.PaymentObj {
	id1, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4d9d")
	id2, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4d9e")

	return []*models.PaymentObj{
		{
			Id:            id1,
			BookingId:     "64b0b9c12f2b5d7c3e6f4d9a",
			Amount:        5000.00,
			Status:        "completed",
			PaymentMethod: "Credit Card",
			TransactionId: "txn_1234567890",
			CreatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            id2,
			BookingId:     "64b0b9c12f2b5d7c3e6f4d9b",
			Amount:        12000.00,
			Status:        "completed",
			PaymentMethod: "Bank Transfer",
			TransactionId: "txn_0987654321",
			CreatedAt:     time.Now().Format(time.RFC3339),
		},
	}
}
