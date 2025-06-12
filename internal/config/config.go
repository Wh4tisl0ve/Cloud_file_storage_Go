package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// todo конфиг для S3

type Config struct {
	Env           string
	Server        HTTPServer
	DB            DB
	Сache         Cache
	ObjectStorage S3
}

type HTTPServer struct {
	Host string
	Port string
}

type DB struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
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

func MustLoad() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getRequiredParam := func(key string) string {
		val := os.Getenv(key)
		if val == "" {
			log.Fatalf("Required environment variable %s not set", key)
		}
		return val
	}

	server := HTTPServer{
		Host: getRequiredParam("SERVER_HOST"),
		Port: getRequiredParam("SERVER_PORT"),
	}

	db := DB{
		Host:     getRequiredParam("DB_HOST"),
		Port:     getRequiredParam("DB_PORT"),
		Name:     getRequiredParam("DB_NAME"),
		Username: getRequiredParam("DB_USERNAME"),
		Password: getRequiredParam("DB_PASSWORD"),
	}

	config := Config{
		Env:    os.Getenv("ENV"),
		Server: server,
		DB:     db,
		//Cache: cache,
		// S3: objectStorage,
	}

	return config
}
