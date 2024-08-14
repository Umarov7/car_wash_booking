package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewService struct {
	Name          string  `bson:"name"`
	Description   string  `bson:"description"`
	Price         float32 `bson:"price"`
	Duration      int32   `bson:"duration"`
	TotalBookings int32   `bson:"total_bookings"`
	CreatedAt     string  `bson:"created_at"`
	UpdatedAt     string  `bson:"updated_at"`
}

type NewServiceData struct {
	Id          string  `bson:"_id"`
	Name        string  `bson:"name"`
	Description string  `bson:"description"`
	Price       float32 `bson:"price"`
	Duration    int32   `bson:"duration"`
	UpdatedAt   string  `bson:"updated_at"`
}

type FilterService struct {
	Name     string  `bson:"name"`
	Price    float32 `bson:"price"`
	Duration int32   `bson:"duration"`
}

type Service struct {
	Id            string  `bson:"_id"`
	Name          string  `bson:"name"`
	Description   string  `bson:"description"`
	Price         float32 `bson:"price"`
	Duration      int32   `bson:"duration"`
	TotalBookings int32   `bson:"total_bookings"`
	CreatedAt     string  `bson:"created_at"`
	UpdatedAt     string  `bson:"updated_at"`
}

type Services struct {
	Services []*Service `bson:"services"`
}

type ServiceObj struct {
	Id            primitive.ObjectID `bson:"_id"`
	Name          string             `bson:"name"`
	Description   string             `bson:"description"`
	Price         float32            `bson:"price"`
	Duration      int32              `bson:"duration"`
	TotalBookings int32              `bson:"total_bookings"`
	CreatedAt     string             `bson:"created_at"`
	UpdatedAt     string             `bson:"updated_at"`
}
