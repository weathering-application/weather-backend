package services

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RedisService struct {
	redisClient *redis.Client
}

func NewRedisService(redisClient *redis.Client) *RedisService {
	return &RedisService{redisClient: redisClient}
}

func (s *RedisService) Set(key string, value string, expiration time.Duration) error {
	ctx := context.Background()
	err := s.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to set key: %v", err)
	}
	return nil
}

func (s *RedisService) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := s.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", status.Errorf(codes.Internal, "failed to get key: %v", err)
	}
	return val, nil
}

func (s *RedisService) Del(key string) error {
	ctx := context.Background()
	err := s.redisClient.Del(ctx, key).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to delete key: %v", err)
	}
	return nil
}
