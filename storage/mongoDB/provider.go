package mongodb

import (
	"booking-service/models"
	"booking-service/storage"
	"context"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProviderRepo struct {
	col *mongo.Collection
}

func NewProviderRepo(db *mongo.Database) storage.IProviderStorage {
	return &ProviderRepo{col: db.Collection("providers")}
}

func (r *ProviderRepo) Create(ctx context.Context, req *models.NewProvider) (string, error) {
	res, err := r.col.InsertOne(ctx, req)
	if err != nil {
		return "", errors.Wrap(err, "query execution failed")
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("id conversion failed")
	}
	return id.Hex(), nil
}

func (r *ProviderRepo) Search(ctx context.Context, req *models.FilterProvider) (*models.Providers, error) {
	filter := bson.M{}
	opts := options.Find()

	if req.CompanyName != "" {
		filter["company_name"] = bson.M{"$regex": req.CompanyName, "$options": "i"}
	}
	if req.AverageRating > 0 {
		filter["average_rating"] = bson.M{"$gte": req.AverageRating}
	}
	if req.CreatedAt != "" {
		filter["created_at"] = bson.M{"$gte": req.CreatedAt}
	}

	cur, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "query execution failed")
	}
	defer cur.Close(ctx)

	var providers []*models.Provider
	for cur.Next(ctx) {
		var pr models.Provider
		if err := cur.Decode(&pr); err != nil {
			return nil, errors.Wrap(err, "decoding failed")
		}
		providers = append(providers, &pr)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return &models.Providers{Providers: providers}, nil
}

func (r *ProviderRepo) UpdateRating(ctx context.Context, id string, rating float32) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	filter := bson.M{"_id": objId}

	var pr models.Provider
	err = r.col.FindOne(ctx, filter).Decode(pr)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}

	newRating := (pr.AverageRating + rating) / 2

	update := bson.M{"$set": bson.M{"average_rating": newRating}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}
