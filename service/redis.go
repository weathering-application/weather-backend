package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/weather-app/monad"
	utils "github.com/weather-app/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CacheService interface {
	Set(ctx context.Context, key string, value []byte, expiration time.Duration) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

type cacheServiceImpl struct {
	redisClient *redis.Client
}

func NewCacheService(redisClient *redis.Client) CacheService {
	return &cacheServiceImpl{redisClient: redisClient}
}

func GetCachedData[T any](ctx context.Context, cacheService CacheService, key string) monad.IO[T] {
	return monad.IO[T]{Run: func() (T, error) {
		cachedData, err := cacheService.Get(ctx, key)
		if err != nil {
			var zero T
			return zero, err
		}

		result, err := utils.UnmarshalData[T]([]byte(cachedData))
		if err != nil {
			var zero T
			return zero, err
		}

		return result, nil
	}}
}

func CacheData(ctx context.Context, cacheService CacheService, key string, data any, expiration time.Duration) monad.IO[error] {
	return monad.IO[error]{Run: func() (error, error) {
		marshaledData, err := utils.MarshalData(data)
		if err != nil {
			return nil, err
		}

		err = cacheService.Set(ctx, key, marshaledData, expiration)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}}
}

func (s *cacheServiceImpl) Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {
	err := s.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to set key: %v", err)
	}
	return nil
}

func (s *cacheServiceImpl) Get(ctx context.Context, key string) (string, error) {
	val, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to get key: %v", err)
	}
	return val, nil
}

func (s *cacheServiceImpl) Del(ctx context.Context, key string) error {
	err := s.redisClient.Del(ctx, key).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to delete key: %v", err)
	}
	return nil
}
