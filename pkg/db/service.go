package db

import (
	"booking-service/models"
	"time"
)

func serviceData() []*models.Service {
	return []*models.Service{
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d94",
			Name:          "Custom Software Development",
			Description:   "Tailor-made software solutions to meet specific business needs.",
			Price:         5000.00,
			Duration:      43200, // 30 days * 24 hours/day * 60 minutes/hour
			TotalBookings: 15,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d95",
			Name:          "IT Consultation",
			Description:   "Expert advice on IT infrastructure and strategy.",
			Price:         200.00,
			Duration:      120, // 2 hours * 60 minutes/hour
			TotalBookings: 40,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d96",
			Name:          "Solar Panel Installation",
			Description:   "Complete installation of solar panels for residential and commercial use.",
			Price:         12000.00,
			Duration:      10080, // 7 days * 24 hours/day * 60 minutes/hour
			TotalBookings: 25,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d97",
			Name:          "Energy Consulting",
			Description:   "Consultation on optimizing energy usage and implementing green technologies.",
			Price:         300.00,
			Duration:      180, // 3 hours * 60 minutes/hour
			TotalBookings: 30,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d98",
			Name:          "Graphic Design",
			Description:   "Custom graphic design services including branding and marketing materials.",
			Price:         1000.00,
			Duration:      14400, // 10 days * 24 hours/day * 60 minutes/hour
			TotalBookings: 50,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		},
		{
			Id:            "64b0b9c12f2b5d7c3e6f4d99",
			Name:          "Web Design",
			Description:   "Design and development of responsive and engaging websites.",
			Price:         2500.00,
			Duration:      28800, // 20 days * 24 hours/day * 60 minutes/hour
			TotalBookings: 45,
			CreatedAt:     time.Now().Format(time.RFC3339),
			UpdatedAt:     time.Now().Format(time.RFC3339),
		},
	}
}
