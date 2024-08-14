package db

import (
	"booking-service/models"
	"time"
)

func providerData() []*models.Provider {
	return []*models.Provider{
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d91",
			UserId:        "4f7a2b34-0d1c-45e3-8a4f-2c4f6b9c0d5b",
			CompanyName:   "Tech Solutions Inc.",
			Description:   "Providing cutting-edge tech solutions for businesses.",
			Services:      []string{"Software Development", "IT Consulting"},
			Availability:  []string{"Monday to Friday", "9 AM to 6 PM"},
			AverageRating: 4.5,
			Location: models.Location{
				Address:   "123 Tech Lane",
				City:      "Innovate City",
				Country:   "Techland",
				Latitude:  37.7749,
				Longitude: -122.4194,
			},
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d92",
			UserId:        "7d56e4a1-1d62-4b5e-9a7a-28f6f6d2b3c4",
			CompanyName:   "Green Energy Solutions",
			Description:   "Expertise in sustainable energy and green technologies.",
			Services:      []string{"Solar Panel Installation", "Energy Consulting"},
			Availability:  []string{"Monday to Saturday", "8 AM to 5 PM"},
			AverageRating: 4.8,
			Location: models.Location{
				Address:   "456 Greenway Drive",
				City:      "Eco Town",
				Country:   "Greenland",
				Latitude:  40.7128,
				Longitude: -74.0060,
			},
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d93",
			UserId:        "3e7f8a1b-5c6d-4e8a-9a0b-89e4d6f7c2d9",
			CompanyName:   "Creative Designs Studio",
			Description:   "Bringing your creative vision to life with innovative design solutions.",
			Services:      []string{"Graphic Design", "Web Design"},
			Availability:  []string{"Monday to Friday", "10 AM to 4 PM"},
			AverageRating: 4.2,
			Location: models.Location{
				Address:   "789 Design Avenue",
				City:      "Artistic City",
				Country:   "Designland",
				Latitude:  34.0522,
				Longitude: -118.2437,
			},
			CreatedAt: time.Now().Format(time.RFC3339),
			UpdatedAt: time.Now().Format(time.RFC3339),
		},
	}
}
