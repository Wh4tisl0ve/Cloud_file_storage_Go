package app

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/config"
	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/repository"
	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/package/postgres"
)

func Run(cfg *config.Config) {
	logger := setupLogger(cfg.Env)
	dbConfig := cfg.Database

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", dbConfig.Username, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	postgres := postgres.New(dsn)

	err := postgres.Connect()
	if err != nil {
		logger.Error(err.Error())
	} else {
		logger.Info("✅ Connected to PostgreSQL successfully!")
	}
	defer postgres.Close()

	/* --Repository-- */

	uRepo := repository.User{
		Conn: postgres.Conn,
	}

	u, err := uRepo.Find(5)
	if err != nil {
		logger.Error(err.Error())
	}
	fmt.Print(u)

}

func setupLogger(env string) *slog.Logger {
	if env == "prod" {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}
