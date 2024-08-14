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

func (r *ProviderRepo) Get(ctx context.Context, id string) (*models.Provider, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id")
	}

	res := r.col.FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "query execution failed")
	}

	var pr models.Provider
	if err := res.Decode(&pr); err != nil {
		return nil, errors.Wrap(err, "decoding failed")
	}
	return &pr, nil
}

func (r *ProviderRepo) Update(ctx context.Context, req *models.NewProviderData) error {
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	set := bson.M{"updated_at": req.UpdatedAt}
	if req.CompanyName != "" {
		set["company_name"] = req.CompanyName
	}
	if req.Description != "" {
		set["description"] = req.Description
	}
	if req.Services != nil {
		set["services"] = req.Services
	}
	if req.Availability != nil {
		set["availability"] = req.Availability
	}
	if req.AverageRating > 0 {
		set["average_rating"] = req.AverageRating
	}
	if req.Location != (models.Location{}) {
		set["location"] = req.Location
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": set}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}

func (r *ProviderRepo) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	res, err := r.col.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}

	if res.DeletedCount == 0 {
		return errors.New("document not found")
	}

	return nil
}

func (r *ProviderRepo) Fetch(ctx context.Context, page, limit int64) (*models.Providers, error) {
	opts := options.Find()
	opts.SetSkip((page - 1) * limit)
	opts.SetLimit(limit)

	cur, err := r.col.Find(ctx, bson.M{}, opts)
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

func (r *ProviderRepo) Search(ctx context.Context, req *models.FilterProvider) (*models.Providers, error) {
	filter := bson.M{}
	opts := options.Find()

	if req.CompanyName != "" {
		filter["company_name"] = bson.M{"$regex": req.CompanyName, "$options": "i"}
	}
	if req.AverageRating > 0 {
		filter["average_rating"] = bson.M{"$gte": req.AverageRating}
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
