package service

import (
	pb "booking-service/genproto/payments"
	"booking-service/models"
	"booking-service/pkg/logger"
	"booking-service/storage"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type PaymentService struct {
	pb.UnimplementedPaymentsServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewPaymentService(s storage.IStorage) *PaymentService {
	return &PaymentService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.NewPayment) (*pb.CreateResp, error) {
	s.logger.Info("CreatePayment is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	pay := models.NewPayment{
		BookingId:     req.BookingId,
		Amount:        req.Amount,
		Status:        req.Status,
		PaymentMethod: req.PaymentMethod,
		TransactionId: req.TransactionId,
		CreatedAt:     time,
	}

	if req.Status == "completed" {
		err := s.storage.Booking().Update(ctx, &models.NewBookingData{
			Id:         req.BookingId,
			Status:     "completed",
			TotalPrice: req.Amount,
			UpdatedAt:  time,
		})
		if err != nil {
			er := errors.Wrap(err, "failed to update booking")
			s.logger.Error(er.Error())
			return nil, er
		}
	}

	id, err := s.storage.Payment().Create(ctx, &pay)
	if err != nil {
		er := errors.Wrap(err, "failed to create payment")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.CreateResp{Id: id, CreatedAt: time}
	s.logger.Info("CreatePayment is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, req *pb.ID) (*pb.Payment, error) {
	s.logger.Info("GetPayment is invoked", slog.Any("request", req))

	pay, err := s.storage.Payment().Get(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get payment")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.Payment{
		Id:            pay.Id,
		BookingId:     pay.BookingId,
		Amount:        pay.Amount,
		Status:        pay.Status,
		PaymentMethod: pay.PaymentMethod,
		TransactionId: pay.TransactionId,
		CreatedAt:     pay.CreatedAt,
	}

	s.logger.Info("GetPayment is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *PaymentService) ListPayments(ctx context.Context, req *pb.Pagination) (*pb.PaymentsList, error) {
	s.logger.Info("ListPayments is invoked", slog.Any("request", req))

	pay, err := s.storage.Payment().Fetch(ctx, int64(req.Page), int64(req.Limit))
	if err != nil {
		er := errors.Wrap(err, "failed to list payments")
		s.logger.Error(er.Error())
		return nil, er
	}

	var payments []*pb.Payment
	for _, p := range pay.Payments {
		payments = append(payments, &pb.Payment{
			Id:            p.Id,
			BookingId:     p.BookingId,
			Amount:        p.Amount,
			Status:        p.Status,
			PaymentMethod: p.PaymentMethod,
			TransactionId: p.TransactionId,
			CreatedAt:     p.CreatedAt,
		})
	}

	s.logger.Info("ListPayments is completed", slog.Any("response", payments))
	return &pb.PaymentsList{Payments: payments, Page: req.Page, Limit: req.Limit}, nil
}
