package storage

import (
	"booking-service/models"
	"context"
)

type IStorage interface {
	Provider() IProviderStorage
	Service() IServiceStorage
	Booking() IBookingStorage
	Payment() IPaymentStorage
	Review() IReviewStorage
	Notification() INotificationStorage
	Close()
}

type IProviderStorage interface {
	Create(ctx context.Context, req *models.NewProvider) (string, error)
	Search(ctx context.Context, req *models.FilterProvider) (*models.Providers, error)
	UpdateRating(ctx context.Context, id string, rating float32) error
}

type IServiceStorage interface {
	Create(ctx context.Context, req *models.NewService) (string, error)
	Update(ctx context.Context, req *models.NewServiceData) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, page, limit int64) (*models.Services, error)
	Search(ctx context.Context, req *models.FilterService) (*models.Services, error)
}

type IBookingStorage interface {
	Create(ctx context.Context, req *models.NewBooking) (string, error)
	Get(ctx context.Context, id string) (*models.Booking, error)
	Update(ctx context.Context, req *models.NewBookingData) error
	Cancel(ctx context.Context, id string) error
	Fetch(ctx context.Context, page, limit int64) (*models.Bookings, error)
}

type IPaymentStorage interface {
	Create(ctx context.Context, req *models.NewPayment) (string, error)
	Get(ctx context.Context, id string) (*models.Payment, error)
	Fetch(ctx context.Context, page, limit int64) (*models.Payments, error)
}

type IReviewStorage interface {
	Create(ctx context.Context, req *models.NewReview) (string, error)
	Update(ctx context.Context, req *models.NewReviewData) error
	Delete(ctx context.Context, id string) error
	Fetch(ctx context.Context, page, limit int64) (*models.Reviews, error)
}

type INotificationStorage interface {
	Create(ctx context.Context, req *models.NewNotification) (string, error)
	Get(ctx context.Context, id string) (*models.Notification, error)
}
