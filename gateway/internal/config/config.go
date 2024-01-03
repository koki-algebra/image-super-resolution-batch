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
	ServerPort         int    `env:"SERVER_PORT" envDefault:"80"`
	ServerAllowOrigins string `env:"SERVER_ALLOW_ORIGINS" envDefault:"http://localhost:8000"`
	// Object storage
	StorageEndpoint                   string `env:"STORAGE_ENDPOINT"`
	StorageBucket                     string `env:"STORAGE_BUCKET" envDefault:"image-super-resolution-batch"`
	StorageUploadImagePrefix          string `env:"STORAGE_UPLOAD_IMAGE_PREFIX" envDefault:"upload_images"`
	StorageSuperResolutionImagePrefix string `env:"STORAGE_SUPER_RESOLUTION_IMAGE_PREFIX" envDefault:"super_resolution_images"`
	// Message queue
	MQHost      string `env:"MQ_HOST" envDefault:"mq"`
	MQPort      int    `env:"MQ_PORT" envDefault:"5672"`
	MQUser      string `env:"MQ_USER" envDefault:"admin"`
	MQPassword  string `env:"MQ_PASSWORD" envDefault:"password"`
	MQQueueName string `env:"MQ_QUEUE_NAME" envDefault:"task_queue"`
}

func New() (*Config, error) {
	cfg := new(Config)
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
