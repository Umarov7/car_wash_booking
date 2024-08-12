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

type ServiceRepo struct {
	col *mongo.Collection
}

func NewServiceRepo(db *mongo.Database) storage.IServiceStorage {
	return &ServiceRepo{col: db.Collection("services")}
}

func (r *ServiceRepo) Create(ctx context.Context, req *models.NewService) (string, error) {
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

func (r *ServiceRepo) Update(ctx context.Context, req *models.NewServiceData) error {
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{
		"name": req.Name, "description": req.Description,
		"price": req.Price, "duration": req.Duration, "updated_at": req.UpdatedAt,
	}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}

func (r *ServiceRepo) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	_, err = r.col.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}

func (r *ServiceRepo) Fetch(ctx context.Context, page, limit int64) (*models.Services, error) {
	opts := options.Find()
	opts.SetSkip((page - 1) * limit)
	opts.SetLimit(limit)

	cur, err := r.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "query execution failed")
	}
	defer cur.Close(ctx)

	var services []*models.Service
	for cur.Next(ctx) {
		var sv models.Service
		if err := cur.Decode(&sv); err != nil {
			return nil, errors.Wrap(err, "decoding failed")
		}
		services = append(services, &sv)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return &models.Services{Services: services}, nil
}

func (r *ServiceRepo) Search(ctx context.Context, req *models.FilterService) (*models.Services, error) {
	filter := bson.M{}
	opts := options.Find()

	if req.Name != "" {
		filter["name"] = bson.M{"$regex": req.Name, "$options": "i"}
	}
	if req.Price > 0 {
		filter["price"] = bson.M{"$gte": req.Price}
	}
	if req.Duration > 0 {
		filter["duration"] = bson.M{"$gte": req.Duration}
	}
	if req.CreatedAt != "" {
		filter["created_at"] = bson.M{"$gte": req.CreatedAt}
	}

	cur, err := r.col.Find(ctx, filter, opts)
	if err != nil {
		return nil, errors.Wrap(err, "query execution failed")
	}
	defer cur.Close(ctx)

	var services []*models.Service
	for cur.Next(ctx) {
		var sv models.Service
		if err := cur.Decode(&sv); err != nil {
			return nil, errors.Wrap(err, "decoding failed")
		}
		services = append(services, &sv)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return &models.Services{Services: services}, nil
}
