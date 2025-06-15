package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Empty command(expected: up, down)")
	}
	cmd := os.Args[1]

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DB.Username, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.Name)

	m, err := migrate.New("file://migrations", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	switch cmd {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal(err)
		}
		log.Print("Migration up successfully")
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal(err)
		}
		log.Print("Migration down successfully")
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}
