package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	config := config.MustLoad()
	dbConfig := config.Database

	dbConnStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	m, err := migrate.New("file://migrations", dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Close()

	if len(os.Args) < 2 {
		log.Fatal("Empty command arguments")
	}

	cmd := strings.Trim(os.Args[1], " ")

	switch cmd {
	case "up":
		if err := m.Up(); err != nil {
			log.Fatal("Error up migrations: ", err)
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal("Error down migrations: ",err)
		}
	default:
		log.Fatal("Unknown command")
	}
}
