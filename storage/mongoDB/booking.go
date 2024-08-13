package mongodb

import (
	"booking-service/models"
	"booking-service/storage"
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookingRepo struct {
	col *mongo.Collection
}

func NewBookingRepo(db *mongo.Database) storage.IBookingStorage {
	return &BookingRepo{col: db.Collection("bookings")}
}

func (r *BookingRepo) Create(ctx context.Context, req *models.NewBooking) (string, error) {
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

func (r *BookingRepo) Get(ctx context.Context, id string) (*models.Booking, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.Wrap(err, "invalid id")
	}

	res := r.col.FindOne(ctx, bson.M{"_id": objId})
	if res.Err() != nil {
		return nil, errors.Wrap(res.Err(), "query execution failed")
	}

	var b models.Booking
	if err := res.Decode(&b); err != nil {
		return nil, errors.Wrap(err, "decoding failed")
	}

	return &b, nil
}

func (r *BookingRepo) Update(ctx context.Context, req *models.NewBookingData) error {
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	set := bson.M{"updated_at": req.UpdatedAt}
	if req.Status != "" {
		set["status"] = req.Status
	}
	if req.ScheduledAt != "" {
		set["scheduled_at"] = req.ScheduledAt
	}
	if req.Location.Address != "" {
		set["location.address"] = req.Location.Address
	}
	if req.Location.City != "" {
		set["location.city"] = req.Location.City
	}
	if req.Location.Country != "" {
		set["location.country"] = req.Location.Country
	}
	if req.Location.Latitude != 0 {
		set["location.latitude"] = req.Location.Latitude
	}
	if req.Location.Longitude != 0 {
		set["location.longitude"] = req.Location.Longitude
	}
	if req.TotalPrice > 0 {
		set["total_price"] = req.TotalPrice
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": set}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}

func (r *BookingRepo) Cancel(ctx context.Context, id string) error {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.Wrap(err, "invalid id")
	}

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{
		"status":     "cancelled",
		"updated_at": time.Now().Format(time.RFC3339),
	}}

	_, err = r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return errors.Wrap(err, "query execution failed")
	}
	return nil
}

func (r *BookingRepo) Fetch(ctx context.Context, page, limit int64) (*models.Bookings, error) {
	opts := options.Find()
	opts.SetSkip((page - 1) * limit)
	opts.SetLimit(limit)

	cur, err := r.col.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, errors.Wrap(err, "query execution failed")
	}
	defer cur.Close(ctx)

	var bookings []*models.Booking
	for cur.Next(ctx) {
		var b models.Booking
		if err := cur.Decode(&b); err != nil {
			return nil, errors.Wrap(err, "decoding failed")
		}
		bookings = append(bookings, &b)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}
	return &models.Bookings{Bookings: bookings}, nil
}
