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

func insertProviders(ctx context.Context, c *mongo.Collection, data []*models.ProviderObj) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "providers")
}

func insertServices(ctx context.Context, c *mongo.Collection, data []*models.ServiceObj) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "services")
}

func insertBookings(ctx context.Context, c *mongo.Collection, data []*models.BookingObj) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "bookings")
}

func insertPayments(ctx context.Context, c *mongo.Collection, data []*models.PaymentObj) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "payments")
}

func insertReviews(ctx context.Context, c *mongo.Collection, data []*models.ReviewObj) error {
	return insertData(ctx, c, convertToInterfaceSlice(data), "reviews")
}
