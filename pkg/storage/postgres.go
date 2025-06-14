package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	DB *sql.DB
}

func New(dsn string) (*Storage, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed create connect to db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed ping to db: %v", err)
	}

	return &Storage{DB: db}, nil
}

func (s *Storage) Close() error{
	err := s.DB.Close()
	if err != nil {
		return err
	}
	return nil
}
