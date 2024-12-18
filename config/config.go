package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ServerAddress string `env:"TODO_SERVER_ADDRESS" env-default:"localhost:7540"`
	Version       string `env:"TODO_VERSION" env-default:"v1.0"`
	DBFile string `env:"TODO_DBFILE" env-default:"./scheduler.db"`
}

func MustLoad() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}
	return &cfg, nil
}
