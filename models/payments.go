package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NewPayment struct {
	BookingId     string  `bson:"booking_id"`
	Amount        float32 `bson:"amount"`
	Status        string  `bson:"status"`
	PaymentMethod string  `bson:"payment_method"`
	TransactionId string  `bson:"transaction_id"`
	CreatedAt     string  `bson:"created_at"`
}

type Payment struct {
	Id            string  `bson:"_id"`
	BookingId     string  `bson:"booking_id"`
	Amount        float32 `bson:"amount"`
	Status        string  `bson:"status"`
	PaymentMethod string  `bson:"payment_method"`
	TransactionId string  `bson:"transaction_id"`
	CreatedAt     string  `bson:"created_at"`
}

type Payments struct {
	Payments []*Payment `bson:"payments"`
}

type PaymentObj struct {
	Id            primitive.ObjectID `bson:"_id"`
	BookingId     string             `bson:"booking_id"`
	Amount        float32            `bson:"amount"`
	Status        string             `bson:"status"`
	PaymentMethod string             `bson:"payment_method"`
	TransactionId string             `bson:"transaction_id"`
	CreatedAt     string             `bson:"created_at"`
}
