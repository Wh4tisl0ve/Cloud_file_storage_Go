package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func New(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed create connect to db: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed ping to db: %v", err)
	}

	return db, nil
}
