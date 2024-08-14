package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewBooking struct {
	UserId      string   `bson:"user_id"`
	ProviderId  string   `bson:"provider_id"`
	ServiceId   string   `bson:"service_id"`
	Status      string   `bson:"status"`
	ScheduledAt string   `bson:"scheduled_at"`
	Location    Location `bson:"location"`
	TotalPrice  float32  `bson:"total_price"`
	CreatedAt   string   `bson:"created_at"`
	UpdatedAt   string   `bson:"updated_at"`
}

type Booking struct {
	Id          string   `bson:"_id"`
	UserId      string   `bson:"user_id"`
	ProviderId  string   `bson:"provider_id"`
	ServiceId   string   `bson:"service_id"`
	Status      string   `bson:"status"`
	ScheduledAt string   `bson:"scheduled_at"`
	Location    Location `bson:"location"`
	TotalPrice  float32  `bson:"total_price"`
	CreatedAt   string   `bson:"created_at"`
	UpdatedAt   string   `bson:"updated_at"`
}

type NewBookingData struct {
	Id          string   `bson:"_id"`
	Status      string   `bson:"status"`
	ScheduledAt string   `bson:"scheduled_at"`
	Location    Location `bson:"location"`
	TotalPrice  float32  `bson:"total_price"`
	UpdatedAt   string   `bson:"updated_at"`
}

type Bookings struct {
	Bookings []*Booking `bson:"bookings"`
}

type BookingObj struct {
	Id          primitive.ObjectID `bson:"_id"`
	UserId      string             `bson:"user_id"`
	ProviderId  string             `bson:"provider_id"`
	ServiceId   string             `bson:"service_id"`
	Status      string             `bson:"status"`
	ScheduledAt string             `bson:"scheduled_at"`
	Location    Location           `bson:"location"`
	TotalPrice  float32            `bson:"total_price"`
	CreatedAt   string             `bson:"created_at"`
	UpdatedAt   string             `bson:"updated_at"`
}
