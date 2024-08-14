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

type ReviewRepo struct {
	col *mongo.Collection
}

func NewReviewRepo(db *mongo.Database) storage.IReviewStorage {
	return &ReviewRepo{col: db.Collection("reviews")}
}

func (r *ReviewRepo) Create(ctx context.Context, req *models.NewReview) (string, error) {
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

func (r *ReviewRepo) Get(ctx context.Context, id string) (*models.Review, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id")
	}

	res := r.col.FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "query execution failed")
	}

	var rv models.Review
	if err := res.Decode(&rv); err != nil {
		return nil, errors.Wrap(err, "decoding failed")
	}
	return &rv, nil
}

func (r *ReviewRepo) Update(ctx context.Context, req *models.NewReviewData) error {
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{
		"rating": req.Rating, "comment": req.Comment, "updated_at": req.UpdatedAt,
	}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}

func (r *ReviewRepo) Delete(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	filter := bson.M{"_id": objId}
	res, err := r.col.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}

	if res.DeletedCount == 0 {
		return errors.New("document not found")
	}

	return nil
}

func (r *ReviewRepo) Fetch(ctx context.Context, page, limit int64) (*models.Reviews, error) {
	opts := options.Find()
	opts.SetSkip((page - 1) * limit)
	opts.SetLimit(limit)

	cur, err := r.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "query execution failed")
	}
	defer cur.Close(ctx)

	var reviews []*models.Review
	for cur.Next(ctx) {
		var r models.Review
		if err := cur.Decode(&r); err != nil {
			return nil, errors.Wrap(err, "decoding failed")
		}
		reviews = append(reviews, &r)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return &models.Reviews{Reviews: reviews}, nil
}
