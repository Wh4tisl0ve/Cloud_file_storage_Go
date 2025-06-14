package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Env           string `env:"ENV" envDefault:"dev"`
	Server        HTTPServer
	DB            DB
	Cache         Cache
	ObjectStorage S3
}

type HTTPServer struct {
	Host string `env:"SERVER_HOST,required"`
	Port string `env:"SERVER_PORT,required"`
}

type DB struct {
	Host     string `env:"DB_HOST,required"`
	Port     string `env:"DB_PORT,required"`
	Name     string `env:"DB_NAME,required"`
	Username string `env:"DB_USERNAME,required"`
	Password string `env:"DB_PASSWORD,required"`
}

type Cache struct {
	Host     string
	Port     string
	Username string
	Password string
}

type S3 struct {
	Host       string
	Port       string
	BucketName string
	AccessKey  string
	SecretKey  string
}

func MustLoad() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error load .env file: %w", err)
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
