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

func Service() storage.IServiceStorage {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil
	}

	return NewServiceRepo(client.Database("test"))
}

func TestCreateService(t *testing.T) {
	s := Service()

	_, err := s.Create(context.Background(), &models.NewService{
		Name:        "test",
		Description: "test",
		Price:       100,
		Duration:    1,
	})
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateService(t *testing.T) {
	s := Service()

	err := s.Update(context.Background(), &models.NewServiceData{
		Id:          "", // ObjectIdHex
		Name:        "test",
		Description: "test",
		Price:       1,
		Duration:    1,
		UpdatedAt:   time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteService(t *testing.T) {
	s := Service()

	err := s.Delete(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestFetchService(t *testing.T) {
	s := Service()

	_, err := s.Fetch(context.Background(), 1, 10)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchService(t *testing.T) {
	s := Service()

	_, err := s.Search(context.Background(), &models.FilterService{Name: "test"})
	if err != nil {
		t.Error(err)
	}
}
