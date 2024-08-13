package service

import (
	pb "booking-service/genproto/reviews"
	"booking-service/models"
	"booking-service/pkg/logger"
	"booking-service/storage"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type ReviewService struct {
	pb.UnimplementedReviewsServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewReviewService(s storage.IStorage) *ReviewService {
	return &ReviewService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *ReviewService) CreateReview(ctx context.Context, req *pb.NewReview) (*pb.CreateResp, error) {
	s.logger.Info("CreateReview is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	rw := models.NewReview{
		BookingId:  req.BookingId,
		UserId:     req.UserId,
		ProviderId: req.ProviderId,
		Rating:     req.Rating,
		Comment:    req.Comment,
		CreatedAt:  time,
		UpdatedAt:  time,
	}

	id, err := s.storage.Review().Create(ctx, &rw)
	if err != nil {
		er := errors.Wrap(err, "failed to create review")
		s.logger.Error(er.Error())
		return nil, er
	}

	err = s.storage.Provider().UpdateRating(ctx, req.ProviderId, float32(req.Rating))
	if err != nil {
		er := errors.Wrap(err, "failed to update provider rating")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.CreateResp{Id: id, CreatedAt: time}
	s.logger.Info("CreateReview is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ReviewService) GetReview(ctx context.Context, req *pb.ID) (*pb.Review, error) {
	s.logger.Info("GetReview is invoked", slog.Any("request", req))

	rw, err := s.storage.Review().Get(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get review")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.Review{
		Id:         rw.Id,
		BookingId:  rw.BookingId,
		UserId:     rw.UserId,
		ProviderId: rw.ProviderId,
		Rating:     rw.Rating,
		Comment:    rw.Comment,
		CreatedAt:  rw.CreatedAt,
		UpdatedAt:  rw.UpdatedAt,
	}

	s.logger.Info("GetReview is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ReviewService) UpdateReview(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	s.logger.Info("UpdateReview is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	rw := models.NewReviewData{
		Id:        req.Id,
		Rating:    req.Rating,
		Comment:   req.Comment,
		UpdatedAt: time,
	}

	err := s.storage.Review().Update(ctx, &rw)
	if err != nil {
		er := errors.Wrap(err, "failed to update review")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.UpdateResp{UpdatedAt: time}
	s.logger.Info("UpdateReview is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ReviewService) DeleteReview(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.logger.Info("DeleteReview is invoked", slog.Any("request", req))

	err := s.storage.Review().Delete(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete review")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("DeleteReview is completed")
	return &pb.Void{}, nil
}

func (s *ReviewService) ListReviews(ctx context.Context, req *pb.Pagination) (*pb.ReviewsList, error) {
	s.logger.Info("ListReviews is invoked", slog.Any("request", req))

	rws, err := s.storage.Review().Fetch(ctx, int64(req.Page), int64(req.Limit))
	if err != nil {
		er := errors.Wrap(err, "failed to list reviews")
		s.logger.Error(er.Error())
		return nil, er
	}

	var reviews []*pb.Review
	for _, r := range rws.Reviews {
		reviews = append(reviews, &pb.Review{
			Id:         r.Id,
			BookingId:  r.BookingId,
			UserId:     r.UserId,
			ProviderId: r.ProviderId,
			Rating:     r.Rating,
			Comment:    r.Comment,
			CreatedAt:  r.CreatedAt,
			UpdatedAt:  r.UpdatedAt,
		})
	}

	s.logger.Info("ListReviews is completed", slog.Any("response", reviews))
	return &pb.ReviewsList{Reviews: reviews, Page: req.Page, Limit: req.Limit}, nil
}
