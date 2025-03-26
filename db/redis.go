package db

import (
	"github.com/redis/go-redis/v9"
)

// Impure
func ConnectRedis(redisAddr string, redisPassword string, redisDB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
}
