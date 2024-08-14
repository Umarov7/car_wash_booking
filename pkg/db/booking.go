package db

import (
	"booking-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func bookingData() []*models.BookingObj {
	id1, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4d9a")
	id2, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4d9b")
	id3, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4d9c")

	return []*models.BookingObj{
		{
			Id:          id1,
			UserId:      "3c6d3f25-8bde-4f85-9b4f-9d8b8c748b89",
			ProviderId:  "64b0b9c12f2b5d7c3e6f4d91",
			ServiceId:   "64b0b9c12f2b5d7c3e6f4d94",
			Status:      "confirmed",
			ScheduledAt: "2024-09-01T10:00:00Z",
			Location: models.Location{
				Address:   "123 Tech Lane",
				City:      "Innovate City",
				Country:   "Techland",
				Latitude:  37.7749,
				Longitude: -122.4194,
			},
			TotalPrice: 5000.00,
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		},
		{
			Id:          id2,
			UserId:      "1d62f49d-ec4e-4a4c-9d56-28f1f7641d09",
			ProviderId:  "64b0b9c12f2b5d7c3e6f4d92",
			ServiceId:   "64b0b9c12f2b5d7c3e6f4d96",
			Status:      "completed",
			ScheduledAt: "2024-09-10T08:00:00Z",
			Location: models.Location{
				Address:   "456 Greenway Drive",
				City:      "Eco Town",
				Country:   "Greenland",
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			TotalPrice: 12000.00,
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		},
		{
			Id:          id3,
			UserId:      "1b9c3d4a-5f6e-4b7f-9e2a-67e8f5c6d2a7",
			ProviderId:  "64b0b9c12f2b5d7c3e6f4d93",
			ServiceId:   "64b0b9c12f2b5d7c3e6f4d98",
			Status:      "pending",
			ScheduledAt: "2024-09-15T10:00:00Z",
			Location: models.Location{
				Address:   "789 Design Avenue",
				City:      "Artistic City",
				Country:   "Designland",
				Latitude:  34.0522,
				Longitude: -118.2437,
			},
			TotalPrice: 1000.00,
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		},
	}
}
