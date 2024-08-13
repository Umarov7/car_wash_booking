package service

import (
	pb "booking-service/genproto/bookings"
	"booking-service/models"
	"booking-service/pkg/logger"
	"booking-service/storage"
	"booking-service/storage/redis"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type BookingService struct {
	pb.UnimplementedBookingsServer
	storage storage.IStorage
	redis   *redis.Storage
	logger  *slog.Logger
}

func NewBookingService(s storage.IStorage, r *redis.Storage) *BookingService {
	return &BookingService{
		storage: s,
		redis:   r,
		logger:  logger.NewLogger(),
	}
}

func (s *BookingService) CreateBooking(ctx context.Context, req *pb.NewBooking) (*pb.CreateResp, error) {
	s.logger.Info("CreateBooking is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	bk := models.NewBooking{
		UserId:      req.UserId,
		ProviderId:  req.ProviderId,
		ServiceId:   req.ServiceId,
		Status:      req.Status,
		ScheduledAt: req.ScheduledTime,
		Location: models.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		TotalPrice: req.TotalPrice,
		CreatedAt:  time,
		UpdatedAt:  time,
	}

	id, err := s.storage.Booking().Create(ctx, &bk)
	if err != nil {
		er := errors.Wrap(err, "failed to create booking")
		s.logger.Error(er.Error())
		return nil, er
	}

	err = s.storage.Service().IncrementBookings(ctx, bk.ServiceId)
	if err != nil {
		er := errors.Wrap(err, "failed to increment bookings")
		s.logger.Error(er.Error())
		return nil, er
	}

	sv, err := s.storage.Service().GetPopular(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to get popular services")
		s.logger.Error(er.Error())
		return nil, er
	}

	err = s.redis.StoreServices(ctx, sv.Services)
	if err != nil {
		er := errors.Wrap(err, "failed to store popular services in redis")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.CreateResp{Id: id, CreatedAt: time}
	s.logger.Info("CreateBooking is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *BookingService) GetBooking(ctx context.Context, req *pb.ID) (*pb.Booking, error) {
	s.logger.Info("GetBooking is invoked", slog.Any("request", req))

	bk, err := s.storage.Booking().Get(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get booking")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.Booking{
		Id:            bk.Id,
		UserId:        bk.UserId,
		ProviderId:    bk.ProviderId,
		ServiceId:     bk.ServiceId,
		Status:        bk.Status,
		ScheduledTime: bk.ScheduledAt,
		Location: &pb.Location{
			Address:   bk.Location.Address,
			City:      bk.Location.City,
			Country:   bk.Location.Country,
			Latitude:  bk.Location.Latitude,
			Longitude: bk.Location.Longitude,
		},
		TotalPrice: bk.TotalPrice,
		CreatedAt:  bk.CreatedAt,
		UpdatedAt:  bk.UpdatedAt,
	}

	s.logger.Info("GetBooking is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *BookingService) UpdateBooking(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	s.logger.Info("UpdateBooking is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)

	bk := models.NewBookingData{
		Id:          req.Id,
		Status:      req.Status,
		ScheduledAt: req.ScheduledTime,
		Location: models.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		TotalPrice: req.TotalPrice,
		UpdatedAt:  time,
	}

	err := s.storage.Booking().Update(ctx, &bk)
	if err != nil {
		er := errors.Wrap(err, "failed to update booking")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.UpdateResp{UpdatedAt: time}
	s.logger.Info("UpdateBooking is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *BookingService) CancelBooking(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.logger.Info("CancelBooking is invoked", slog.Any("request", req))

	err := s.storage.Booking().Cancel(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to cancel booking")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("CancelBooking is completed")
	return &pb.Void{}, nil
}

func (s *BookingService) ListBookings(ctx context.Context, req *pb.Pagination) (*pb.BookingsList, error) {
	s.logger.Info("ListBookings is invoked", slog.Any("request", req))

	bks, err := s.storage.Booking().Fetch(ctx, int64(req.Page), int64(req.Limit))
	if err != nil {
		er := errors.Wrap(err, "failed to list bookings")
		s.logger.Error(er.Error())
		return nil, er
	}

	var bookings []*pb.Booking
	for _, b := range bks.Bookings {
		bookings = append(bookings, &pb.Booking{
			Id:            b.Id,
			UserId:        b.UserId,
			ProviderId:    b.ProviderId,
			ServiceId:     b.ServiceId,
			Status:        b.Status,
			ScheduledTime: b.ScheduledAt,
			Location: &pb.Location{
				Address:   b.Location.Address,
				City:      b.Location.City,
				Country:   b.Location.Country,
				Latitude:  b.Location.Latitude,
				Longitude: b.Location.Longitude,
			},
			TotalPrice: b.TotalPrice,
			CreatedAt:  b.CreatedAt,
			UpdatedAt:  b.UpdatedAt,
		})
	}

	s.logger.Info("ListBookings is completed", slog.Any("response", bookings))
	return &pb.BookingsList{Bookings: bookings, Page: req.Page, Limit: req.Limit}, nil
}
