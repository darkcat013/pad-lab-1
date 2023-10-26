package cache

import (
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/redis/go-redis/v9"
)

func getCacheStore(cfg config.Config) *redis.Client {
	cacheStore := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConnectionString,
		Password: "",
		DB:       1,
	})

	return cacheStore
}
