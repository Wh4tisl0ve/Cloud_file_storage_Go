package app

import (
	"fmt"
	"log"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/config"
	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/pkg/storage"
)

func Run(cfg *config.Config) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	conn, err := postgres.New(dsn)
	if err != nil {
		log.Fatal(err)
	}
	
	defer conn.Close()

}
