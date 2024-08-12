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

type PaymentRepo struct {
	col *mongo.Collection
}

func NewPaymentRepo(db *mongo.Database) storage.IPaymentStorage {
	return &PaymentRepo{col: db.Collection("payments")}
}

func (r *PaymentRepo) Create(ctx context.Context, req *models.NewPayment) (string, error) {
	res, err := r.col.InsertOne(ctx, req)
	if err != nil {
		return "", err
	}

	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("id conversion failed")
	}
	return id.Hex(), nil
}

func (r *PaymentRepo) Get(ctx context.Context, id string) (*models.Payment, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id")
	}

	res := r.col.FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "query execution failed")
	}

	var p models.Payment
	if err := res.Decode(&p); err != nil {
		return nil, errors.Wrap(err, "decoding failed")
	}
	return &p, nil
}

func (r *PaymentRepo) Fetch(ctx context.Context, page, limit int64) (*models.Payments, error) {
	opts := options.Find()
	opts.SetSkip((page - 1) * limit)
	opts.SetLimit(limit)

	cur, err := r.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "query execution failed")
	}
	defer cur.Close(ctx)

	var payments []*models.Payment
	for cur.Next(ctx) {
		var p models.Payment
		if err := cur.Decode(&p); err != nil {
			return nil, errors.Wrap(err, "decoding failed")
		}
		payments = append(payments, &p)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return &models.Payments{Payments: payments}, nil
}
