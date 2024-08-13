package models

type Location struct {
	Address   string  `bson:"address"`
	City      string  `bson:"city"`
	Country   string  `bson:"country"`
	Latitude  float32 `bson:"latitude"`
	Longitude float32 `bson:"longitude"`
}

type NewProvider struct {
	UserId        string   `bson:"user_id"`
	CompanyName   string   `bson:"company_name"`
	Description   string   `bson:"description"`
	Services      []string `bson:"services"`
	Availability  []string `bson:"availability"`
	AverageRating float32  `bson:"average_rating"`
	Location      Location `bson:"location"`
	CreatedAt     string   `bson:"created_at"`
	UpdatedAt     string   `bson:"updated_at"`
}

type FilterProvider struct {
	CompanyName   string  `bson:"company_name"`
	AverageRating float32 `bson:"average_rating"`
	CreatedAt     string  `bson:"created_at"`
}

type Provider struct {
	Id            string   `bson:"_id"`
	UserId        string   `bson:"user_id"`
	CompanyName   string   `bson:"company_name"`
	Description   string   `bson:"description"`
	Services      []string `bson:"services"`
	Availability  []string `bson:"availability"`
	AverageRating float32  `bson:"average_rating"`
	Location      Location `bson:"location"`
	CreatedAt     string   `bson:"created_at"`
	UpdatedAt     string   `bson:"updated_at"`
}

type Providers struct {
	Providers []*Provider `bson:"providers"`
}
