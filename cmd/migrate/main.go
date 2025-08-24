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
	cfg := config.MustLoad()
	dbCfg := cfg.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbCfg.Username, dbCfg.Password, dbCfg.Host, dbCfg.Port, dbCfg.Name)
	m, err := migrate.New("file://migrations", dsn)
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
		} else{
			log.Println("✅ Migrations up successfully")
		}
	case "down":
		if err := m.Down(); err != nil {
			log.Fatal("Error down migrations: ", err)
		} else{
			log.Println("✅ Migrations down successfully")
		}
	default:
		log.Fatal("Unknown command")
	}
}
