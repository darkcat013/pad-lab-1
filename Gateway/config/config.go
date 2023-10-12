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

	OwnerUrl              string `env:"OWNER_URL" envDefault:"localhost:5171"`
	VeterinaryUrl         string `env:"VET_URL" envDefault:"localhost:5172"`
	RedisConnectionString string `env:"REDIS_CONNECTION_STRING" envDefault:""`
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
