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

func Provider() storage.IProviderStorage {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil
	}

	return NewProviderRepo(client.Database("test"))
}

func TestCreateProvider(t *testing.T) {
	p := Provider()

	_, err := p.Create(context.Background(), &models.NewProvider{
		UserId:        "test",
		CompanyName:   "test",
		Description:   "test",
		Services:      []string{"test"},
		Availability:  []string{"test"},
		AverageRating: 5,
		Location: models.Location{
			Address:   "test",
			City:      "test",
			Country:   "test",
			Latitude:  1.0,
			Longitude: 1.0,
		},
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	})

	if err != nil {
		t.Error(err)
	}
}

func TestGetProvider(t *testing.T) {
	p := Provider()

	_, err := p.Get(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestUpdateProvider(t *testing.T) {
	p := Provider()

	err := p.Update(context.Background(), &models.NewProviderData{
		Id:            "", // ObjectIdHex
		CompanyName:   "test",
		Description:   "test",
		Services:      []string{"testing"},
		AverageRating: 3,
		Location: models.Location{
			Address:   "addres",
			City:      "city",
			Country:   "country",
			Latitude:  1.0,
			Longitude: 1.0,
		},
		UpdatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteProvider(t *testing.T) {
	p := Provider()

	err := p.Delete(context.Background(), "") // ObjectIdHex
	if err != nil {
		t.Error(err)
	}
}

func TestFetchProvider(t *testing.T) {
	p := Provider()

	_, err := p.Fetch(context.Background(), 1, 10)
	if err != nil {
		t.Error(err)
	}
}

func TestSearchProviders(t *testing.T) {
	p := Provider()

	res, err := p.Search(context.Background(), &models.FilterProvider{
		CompanyName: "test",
	})

	if err != nil {
		t.Error(err)
	}

	if len(res.Providers) == 0 {
		t.Error("no providers found")
	}
}
