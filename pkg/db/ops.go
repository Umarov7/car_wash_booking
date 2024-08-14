package db

import (
	"booking-service/models"
	"context"
	"log"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertData(ctx context.Context, c *mongo.Collection, documents []interface{}, collectionName string) error {
	_, err := c.DeleteMany(ctx, bson.M{})
	if err != nil {
		return errors.Wrapf(err, "failed to clear %s", collectionName)
	}

	log.Printf("%s cleared successfully!\n", collectionName)

	_, err = c.InsertMany(ctx, documents)
	if err != nil {
		return errors.Wrapf(err, "failed to insert %s", collectionName)
	}

	log.Printf("sample %s inserted successfully!\n", collectionName)
	return nil
}

func convertToInterfaceSlice[T any](data []T) []interface{} {
	var documents []interface{}
	for _, d := range data {
		documents = append(documents, d)
	}
	return documents
}

func insertProviders(ctx context.Context, c *mongo.Collection, data []*models.Provider) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "providers")
}

func insertServices(ctx context.Context, c *mongo.Collection, data []*models.Service) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "services")
}

func insertBookings(ctx context.Context, c *mongo.Collection, data []*models.Booking) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "bookings")
}

func insertPayments(ctx context.Context, c *mongo.Collection, data []*models.Payment) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "payments")
}

func insertReviews(ctx context.Context, c *mongo.Collection, data []*models.Review) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "reviews")
}
