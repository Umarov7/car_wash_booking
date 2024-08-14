package service

import (
	pb "booking-service/genproto/services"
	"booking-service/models"
	"booking-service/pkg/logger"
	"booking-service/storage"
	"booking-service/storage/redis"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type ServiceService struct {
	pb.UnimplementedServicesServer
	storage storage.IStorage
	redis   *redis.Storage
	logger  *slog.Logger
}

func NewServiceService(s storage.IStorage, r *redis.Storage) *ServiceService {
	return &ServiceService{
		storage: s,
		redis:   r,
		logger:  logger.NewLogger(),
	}
}

func (s *ServiceService) CreateService(ctx context.Context, req *pb.NewService) (*pb.CreateResp, error) {
	s.logger.Info("CreateService is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	sv := models.NewService{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		Duration:      req.Duration,
		TotalBookings: 0, // Default value
		CreatedAt:     time,
		UpdatedAt:     time,
	}

	id, err := s.storage.Service().Create(ctx, &sv)
	if err != nil {
		er := errors.Wrap(err, "failed to create service")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.CreateResp{Id: id, CreatedAt: time}
	s.logger.Info("CreateService is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ServiceService) GetService(ctx context.Context, req *pb.ID) (*pb.Service, error) {
	s.logger.Info("GetService is invoked", slog.Any("request", req))

	sv, err := s.storage.Service().Get(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to get service")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.Service{
		Id:            sv.Id,
		Name:          sv.Name,
		Description:   sv.Description,
		Price:         sv.Price,
		Duration:      sv.Duration,
		TotalBookings: sv.TotalBookings,
		CreatedAt:     sv.CreatedAt,
		UpdatedAt:     sv.UpdatedAt,
	}

	s.logger.Info("GetService is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ServiceService) UpdateService(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	s.logger.Info("UpdateService is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	sv := models.NewServiceData{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Duration:    req.Duration,
		UpdatedAt:   time,
	}

	err := s.storage.Service().Update(ctx, &sv)
	if err != nil {
		er := errors.Wrap(err, "failed to update service")
		s.logger.Error(er.Error())
		return nil, er
	}

	resp := &pb.UpdateResp{UpdatedAt: time}
	s.logger.Info("UpdateService is completed", slog.Any("response", resp))
	return resp, nil
}

func (s *ServiceService) DeleteService(ctx context.Context, req *pb.ID) (*pb.Void, error) {
	s.logger.Info("DeleteService is invoked", slog.Any("request", req))

	err := s.storage.Service().Delete(ctx, req.Id)
	if err != nil {
		er := errors.Wrap(err, "failed to delete service")
		s.logger.Error(er.Error())
		return nil, er
	}

	s.logger.Info("DeleteService is completed")
	return &pb.Void{}, nil
}

func (s *ServiceService) ListServices(ctx context.Context, req *pb.Pagination) (*pb.ServicesList, error) {
	s.logger.Info("ListServices is invoked", slog.Any("request", req))

	sv, err := s.storage.Service().Fetch(ctx, int64(req.Page), int64(req.Limit))
	if err != nil {
		er := errors.Wrap(err, "failed to list services")
		s.logger.Error(er.Error())
		return nil, er
	}

	var services []*pb.Service
	for _, service := range sv.Services {
		services = append(services, &pb.Service{
			Id:            service.Id,
			Name:          service.Name,
			Description:   service.Description,
			Price:         service.Price,
			Duration:      service.Duration,
			TotalBookings: service.TotalBookings,
			CreatedAt:     service.CreatedAt,
			UpdatedAt:     service.UpdatedAt,
		})
	}

	s.logger.Info("ListServices is completed", slog.Any("response", services))
	return &pb.ServicesList{Services: services, Page: req.Page, Limit: req.Limit}, nil
}

func (s *ServiceService) SearchServices(ctx context.Context, req *pb.Filter) (*pb.SearchResp, error) {
	s.logger.Info("SearchServices is invoked", slog.Any("request", req))

	sv, err := s.storage.Service().Search(ctx, &models.FilterService{
		Name:     req.Name,
		Price:    req.Price,
		Duration: req.Duration,
	})
	if err != nil {
		er := errors.Wrap(err, "failed to search services")
		s.logger.Error(er.Error())
		return nil, er
	}

	var services []*pb.Service
	for _, service := range sv.Services {
		services = append(services, &pb.Service{
			Id:            service.Id,
			Name:          service.Name,
			Description:   service.Description,
			Price:         service.Price,
			Duration:      service.Duration,
			TotalBookings: service.TotalBookings,
			CreatedAt:     service.CreatedAt,
			UpdatedAt:     service.UpdatedAt,
		})
	}

	s.logger.Info("SearchServices is completed", slog.Any("response", services))
	return &pb.SearchResp{Services: services}, nil
}

func (s *ServiceService) GetPopularServices(ctx context.Context, req *pb.Void) (*pb.SearchResp, error) {
	s.logger.Info("GetPopularServices is invoked", slog.Any("request", req))

	sv, err := s.redis.GetServices(ctx)
	if err != nil {
		er := errors.Wrap(err, "failed to retrieve popular services from redis")
		s.logger.Error(er.Error())

		sv, err = s.storage.Service().GetPopular(ctx)
		if err != nil {
			er := errors.Wrap(err, "failed to get popular services")
			s.logger.Error(er.Error())
			return nil, er
		}
	} else if len(sv.Services) == 0 {
		s.logger.Info("No popular services found in redis, fetching from database")

		sv, err = s.storage.Service().GetPopular(ctx)
		if err != nil {
			er := errors.Wrap(err, "failed to get popular services")
			s.logger.Error(er.Error())
			return nil, er
		}
	}

	var services []*pb.Service
	for _, service := range sv.Services {
		services = append(services, &pb.Service{
			Id:            service.Id,
			Name:          service.Name,
			Description:   service.Description,
			Price:         service.Price,
			Duration:      service.Duration,
			TotalBookings: service.TotalBookings,
			CreatedAt:     service.CreatedAt,
			UpdatedAt:     service.UpdatedAt,
		})
	}

	s.logger.Info("GetPopularServices is completed", slog.Any("response", services))
	return &pb.SearchResp{Services: services}, nil
}
