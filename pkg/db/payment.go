package db

import (
	"booking-service/models"
	"time"
)

func paymentData() []*models.Payment {
	return []*models.Payment{
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d9d",
			BookingId:     "64b0b9c12f2b5d7c3e6f4d9a",
			Amount:        5000.00,
			Status:        "Completed",
			PaymentMethod: "Credit Card",
			TransactionId: "txn_1234567890",
			CreatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d9e",
			BookingId:     "64b0b9c12f2b5d7c3e6f4d9b",
			Amount:        12000.00,
			Status:        "Completed",
			PaymentMethod: "Bank Transfer",
			TransactionId: "txn_0987654321",
			CreatedAt:     time.Now().Format(time.RFC3339),
		},
	}
}
