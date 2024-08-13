package redis

import (
	"booking-service/config"
	"booking-service/models"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Storage struct {
	config *config.Config
	db     *redis.Client
}

func ConnectDB(cfg *config.Config) (*Storage, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.REDIS_ADDRESS,
		Password: cfg.REDIS_PASSWORD,
		DB:       cfg.REDIS_DB,
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrap(err, "failed to connect to redis")
	}

	return &Storage{config: cfg, db: rdb}, nil
}

func (r *Storage) StoreServices(ctx context.Context, services []*models.Service) error {
	for _, service := range services {
		serviceJSON, err := json.Marshal(service)
		if err != nil {
			return errors.Wrap(err, "failed to serialize service object")
		}

		err = r.db.HSet(ctx, r.config.REDIS_KEY, service.Id, serviceJSON).Err()
		if err != nil {
			return errors.Wrap(err, "failed to store service")
		}

	}

	err := r.db.Expire(ctx, r.config.REDIS_KEY, time.Minute*10).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set expiration time")
	}

	return nil
}

func (r *Storage) GetServices(ctx context.Context) (*models.Services, error) {
	services, err := r.db.HGetAll(ctx, r.config.REDIS_KEY).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get services")
	}

	var res []*models.Service

	for _, s := range services {
		var sv models.Service
		err := json.Unmarshal([]byte(s), &sv)
		if err != nil {
			return nil, errors.Wrap(err, "failed to deserialize service object")
		}
		res = append(res, &sv)
	}

	return &models.Services{Services: res}, nil
}

func (r *Storage) Close() {
	if err := r.db.Close(); err != nil {
		log.Printf("error while disconnecting from redis: %v", err)
	}
}
