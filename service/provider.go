package service

import (
	pb "booking-service/genproto/providers"
	"booking-service/models"
	"booking-service/pkg/logger"
	"booking-service/storage"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type ProviderService struct {
	pb.UnimplementedProvidersServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewProviderService(s storage.IStorage) *ProviderService {
	return &ProviderService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *ProviderService) CreateProvider(ctx context.Context, req *pb.NewProvider) (*pb.CreateResp, error) {
	s.logger.Info("CreateProvider is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	pr := models.NewProvider{
		UserId:        req.UserId,
		CompanyName:   req.CompanyName,
		Description:   req.Description,
		Services:      req.Services,
		Availability:  req.Availability,
		AverageRating: req.AverageRating,
		Location: models.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		CreatedAt: time,
		UpdatedAt: time,
	}

	id, err := s.storage.Provider().Create(ctx, &pr)
	if err != nil {
		er := errors.Wrap(err, "failed to create provider")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.CreateResp{Id: id, CreatedAt: time}
	s.logger.Info("CreateProvider is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ProviderService) GetProvider(ctx context.Context, req *pb.ID) (*pb.Provider, error) {
	s.logger.Info("GetProvider is invoked", slog.Any("request", req))

	pr, err := s.storage.Provider().Get(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get provider")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.Provider{
		Id:            pr.Id,
		UserId:        pr.UserId,
		CompanyName:   pr.CompanyName,
		Description:   pr.Description,
		Services:      pr.Services,
		Availability:  pr.Availability,
		AverageRating: pr.AverageRating,
		Location: &pb.Location{
			Address:   pr.Location.Address,
			City:      pr.Location.City,
			Country:   pr.Location.Country,
			Latitude:  pr.Location.Latitude,
			Longitude: pr.Location.Longitude,
		},
		CreatedAt: pr.CreatedAt,
		UpdatedAt: pr.UpdatedAt,
	}

	s.logger.Info("GetProvider is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ProviderService) UpdateProvider(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	s.logger.Info("UpdateProvider is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	pr := models.NewProviderData{
		Id:            req.Id,
		CompanyName:   req.CompanyName,
		Description:   req.Description,
		Services:      req.Services,
		Availability:  req.Availability,
		AverageRating: req.AverageRating,
		Location: models.Location{
			Address:   req.Location.Address,
			City:      req.Location.City,
			Country:   req.Location.Country,
			Latitude:  req.Location.Latitude,
			Longitude: req.Location.Longitude,
		},
		UpdatedAt: time,
	}

	err := s.storage.Provider().Update(ctx, &pr)
	if err != nil {
		er := errors.Wrap(err, "failed to update provider")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.UpdateResp{UpdatedAt: time}
	s.logger.Info("UpdateProvider is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ProviderService) DeleteProvider(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.logger.Info("DeleteProvider is invoked", slog.Any("request", req))

	err := s.storage.Provider().Delete(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete provider")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("DeleteProvider is completed")
	return &pb.Void{}, nil
}

func (s *ProviderService) ListProviders(ctx context.Context, req *pb.Pagination) (*pb.ProvidersList, error) {
	s.logger.Info("ListProviders is invoked", slog.Any("request", req))

	resp, err := s.storage.Provider().Fetch(ctx, int64(req.Page), int64(req.Limit))
	if err != nil {
		er := errors.Wrap(err, "failed to list providers")
		s.logger.Error(er.Error())
		return nil, er
	}

	var providers []*pb.Provider
	for _, p := range resp.Providers {
		providers = append(providers, &pb.Provider{
			Id:            p.Id,
			UserId:        p.UserId,
			CompanyName:   p.CompanyName,
			Description:   p.Description,
			Services:      p.Services,
			Availability:  p.Availability,
			AverageRating: p.AverageRating,
			Location: &pb.Location{
				Address:   p.Location.Address,
				City:      p.Location.City,
				Country:   p.Location.Country,
				Latitude:  p.Location.Latitude,
				Longitude: p.Location.Longitude,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	s.logger.Info("ListProviders is completed", slog.Any("response", providers))
	return &pb.ProvidersList{Providers: providers, Page: req.Page, Limit: req.Limit}, nil
}

func (s *ProviderService) SearchProviders(ctx context.Context, req *pb.Filter) (*pb.SearchResp, error) {
	s.logger.Info("SearchProviders is invoked", slog.Any("request", req))

	resp, err := s.storage.Provider().Search(ctx, &models.FilterProvider{
		CompanyName:   req.CompanyName,
		AverageRating: req.AverageRating,
		CreatedAt:     req.CreatedAt,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to find providers")
		s.logger.Error(er.Error())
		return nil, er
	}

	var providers []*pb.Provider
	for _, p := range resp.Providers {
		providers = append(providers, &pb.Provider{
			Id:            p.Id,
			UserId:        p.UserId,
			CompanyName:   p.CompanyName,
			Description:   p.Description,
			Services:      p.Services,
			Availability:  p.Availability,
			AverageRating: p.AverageRating,
			Location: &pb.Location{
				Address:   p.Location.Address,
				City:      p.Location.City,
				Country:   p.Location.Country,
				Latitude:  p.Location.Latitude,
				Longitude: p.Location.Longitude,
			},
			CreatedAt: p.CreatedAt,
			UpdatedAt: p.UpdatedAt,
		})
	}

	s.logger.Info("SearchProviders is completed", slog.Any("response", providers))
	return &pb.SearchResp{Providers: providers}, nil
}
