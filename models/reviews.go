package models

type NewReview struct {
	BookingId  string `bson:"booking_id"`
	UserId     string `bson:"user_id"`
	ProviderId string `bson:"provider_id"`
	Rating     int32  `bson:"rating"`
	Comment    string `bson:"comment"`
	CreatedAt  string `bson:"created_at"`
	UpdatedAt  string `bson:"updated_at"`
}

type NewReviewData struct {
	Id        string `bson:"id"`
	Rating    int32  `bson:"rating"`
	Comment   string `bson:"comment"`
	UpdatedAt string `bson:"updated_at"`
}

type Review struct {
	Id         string `bson:"id"`
	BookingId  string `bson:"booking_id"`
	UserId     string `bson:"user_id"`
	ProviderId string `bson:"provider_id"`
	Rating     int32  `bson:"rating"`
	Comment    string `bson:"comment"`
	CreatedAt  string `bson:"created_at"`
	UpdatedAt  string `bson:"updated_at"`
}

type Reviews struct {
	Reviews []*Review `bson:"reviews"`
}
