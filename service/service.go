package service

import (
	pb "booking-service/genproto/services"
	"booking-service/models"
	"booking-service/pkg/logger"
	"booking-service/storage"
	"context"
	"log/slog"
	"time"

	"github.com/pkg/errors"
)

type ServiceService struct {
	pb.UnimplementedServicesServer
	storage storage.IStorage
	logger  *slog.Logger
}

func NewServiceService(s storage.IStorage) *ServiceService {
	return &ServiceService{
		storage: s,
		logger:  logger.NewLogger(),
	}
}

func (s *ServiceService) CreateService(ctx context.Context, req *pb.NewService) (*pb.CreateResp, error) {
	s.logger.Info("CreateService is invoked", slog.Any("request", req))

	time := time.Now().Format(time.RFC3339)
	sv := models.NewService{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Duration:    req.Duration,
		CreatedAt:   time,
		UpdatedAt:   time,
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
			Id:          service.Id,
			Name:        service.Name,
			Description: service.Description,
			Price:       service.Price,
			Duration:    service.Duration,
			CreatedAt:   service.CreatedAt,
			UpdatedAt:   service.UpdatedAt,
		})
	}

	s.logger.Info("ListServices is completed", slog.Any("response", services))
	return &pb.ServicesList{Services: services, Page: req.Page, Limit: req.Limit}, nil
}
