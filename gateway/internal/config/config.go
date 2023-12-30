package config

import "github.com/caarlos0/env"

type Config struct {
	// Database
	DBHost     string `env:"DB_HOST" envDefault:"db"`
	DBPort     int    `env:"DB_PORT" envDefault:"5432"`
	DBDatabase string `env:"DB_DATABASE" envDefault:"app"`
	DBUser     string `env:"DB_USER" envDefault:"postgres"`
	DBPassword string `env:"DB_PASSWORD" envDefault:"password"`
	// Server
	ServerPort int `env:"SERVER_PORT" envDefault:"80"`
	// Object storage
	UploadImagePrefix          string `env:"UPLOAD_IMAGE_PREFIX" envDefault:"upload_images"`
	SuperResolutionImagePrefix string `env:"SUPER_RESOLUTION_IMAGE_PREFIX" envDefault:"super_resolution_images"`
}

func New() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
