package cache

import (
	"github.com/darkcat013/pad-lab-1/Gateway/config"
	"github.com/redis/go-redis/v9"
)

func GetCircuitBreakerStore(cfg config.Config) *redis.Client {
	circuitBreakerStore := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisConnectionString,
		Password: "",
		DB:       2,
	})

	return circuitBreakerStore
}
