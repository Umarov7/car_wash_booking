package mongodb

import (
	"booking-service/config"
	"booking-service/storage"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	db *mongo.Database
}

func ConnectDB(cfg *config.Config) (storage.IStorage, error) {
	opts := options.Client().ApplyURI(cfg.DB_URI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{db: client.Database(cfg.DB_NAME)}, nil
}

func (m *MongoDB) Close() {
	if err := m.db.Client().Disconnect(context.Background()); err != nil {
		log.Fatalf("error while disconnecting from mongodb: %v", err)
	}
}

func (m *MongoDB) Provider() storage.IProviderStorage {
	return NewProviderRepo(m.db)
}

func (m *MongoDB) Service() storage.IServiceStorage {
	return NewServiceRepo(m.db)
}

func (m *MongoDB) Booking() storage.IBookingStorage {
	return NewBookingRepo(m.db)
}

func (m *MongoDB) Payment() storage.IPaymentStorage {
	return NewPaymentRepo(m.db)
}

func (m *MongoDB) Review() storage.IReviewStorage {
	return NewReviewRepo(m.db)
}

func (m *MongoDB) Notification() storage.INotificationStorage {
	return NewNotificationRepo(m.db)
}
