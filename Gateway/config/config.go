package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Host        string `env:"HOST" envDefault:"localhost"`
	Port        string `env:"SERVER_PORT" envDefault:"8080"`
	AllowOrigin string `env:"ALLOW_ORIGIN" envDefault:"*"`

	OwnerUrl              string `env:"OWNER_URL" envDefault:"localhost:32769"`
	VeterinaryUrl         string `env:"VET_URL" envDefault:"localhost:32771"`
	RedisConnectionString string `env:"REDIS_CONNECTION_STRING" envDefault:"localhost:6379"`
	RedisRateLimitDb      int    `env:"REDIS_RATE_LIMIT_DB" envDefault:"0"`
	RedisCacheDb          int    `env:"REDIS_CACHE_DB" envDefault:"1"`
}

func InitConfig() (Config, error) {
	var cfg Config

	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("Failed to load .env file.")
	}

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
