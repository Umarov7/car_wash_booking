package db

import (
	"booking-service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func reviewData() []*models.ReviewObj {
	id1, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4d9f")
	id2, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4da0")
	id3, _ := primitive.ObjectIDFromHex("64b0b9c12f2b5d7c3e6f4da1")

	return []*models.ReviewObj{
		{
			Id:         id1,
			BookingId:  "64b0b9c12f2b5d7c3e6f4d9a",
			UserId:     "3c6d3f25-8bde-4f85-9b4f-9d8b8c748b89",
			ProviderId: "64b0b9c12f2b5d7c3e6f4d91",
			Rating:     5,
			Comment:    "Exceptional service and top-notch software development. Highly recommend!",
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		},
		{
			Id:         id2,
			BookingId:  "64b0b9c12f2b5d7c3e6f4d9b",
			UserId:     "1d62f49d-ec4e-4a4c-9d56-28f1f7641d09",
			ProviderId: "64b0b9c12f2b5d7c3e6f4d92",
			Rating:     4,
			Comment:    "Great installation, but the project took a bit longer than expected.",
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		},
		{
			Id:         id3,
			BookingId:  "64b0b9c12f2b5d7c3e6f4d9c",
			UserId:     "1b9c3d4a-5f6e-4b7f-9e2a-67e8f5c6d2a7",
			ProviderId: "64b0b9c12f2b5d7c3e6f4d93",
			Rating:     3,
			Comment:    "Design was okay, but communication could be improved.",
			CreatedAt:  time.Now().Format(time.RFC3339),
			UpdatedAt:  time.Now().Format(time.RFC3339),
		},
	}
}
