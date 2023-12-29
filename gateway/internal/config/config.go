package config

import "github.com/caarlos0/env"

type Config struct {
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBDatabase string `env:"DB_DATABASE" envDefault:"app"`
	DBUser     string `env:"DB_USER" envDefault:"user"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"password"`
}

func New() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
