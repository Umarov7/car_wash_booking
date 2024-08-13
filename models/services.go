package models

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
	Name      string  `bson:"name"`
	Price     float32 `bson:"price"`
	Duration  int32   `bson:"duration"`
	CreatedAt string  `bson:"created_at"`
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
