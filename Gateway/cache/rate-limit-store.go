package cache

import (
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/redis/go-redis/v9"
)

func GetRateLimitStore(cfg config.Config) *redis.Client {
	rateLimitStore := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConnectionString,
		Password: "",
		DB:       cfg.RedisRateLimitDb,
	})

	return rateLimitStore
}
