package db

import (
	"booking-service/config"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SeedData(ctx context.Context, cfg *config.Config) error {
	opts := options.Client().ApplyURI(cfg.DB_URI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return errors.Wrap(err, "error connecting to database")
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return errors.Wrap(err, "error pinging database")
	}

	providers := client.Database(cfg.DB_NAME).Collection("providers")
	services := client.Database(cfg.DB_NAME).Collection("services")
	bookings := client.Database(cfg.DB_NAME).Collection("bookings")
	payments := client.Database(cfg.DB_NAME).Collection("payments")
	reviews := client.Database(cfg.DB_NAME).Collection("reviews")

	if err := insertProviders(ctx, providers, providerData()); err != nil {
		return errors.Wrap(err, "error inserting providers")
	}

	if err := insertServices(ctx, services, serviceData()); err != nil {
		return errors.Wrap(err, "error inserting services")
	}

	if err := insertBookings(ctx, bookings, bookingData()); err != nil {
		return errors.Wrap(err, "error inserting bookings")
	}

	if err := insertPayments(ctx, payments, paymentData()); err != nil {
		return errors.Wrap(err, "error inserting payments")
	}

	if err := insertReviews(ctx, reviews, reviewData()); err != nil {
		return errors.Wrap(err, "error inserting reviews")
	}

	return nil
}
