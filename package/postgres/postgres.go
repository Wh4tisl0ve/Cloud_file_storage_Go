package postgres

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgreSQL struct {
	dsn  string
	Conn *sql.DB
}

func New(dsn string) *PostgreSQL {
	return &PostgreSQL{
		dsn: dsn,
	}
}

func (p *PostgreSQL) Connect() error {
	conn, err := sql.Open("pgx", p.dsn)
	if err != nil {
		return err
	}

	err = conn.Ping()
	if err != nil {
		return err
	}

	p.Conn = conn

	return nil
}

func (p *PostgreSQL) Close() error {
	err := p.Conn.Close()

	if err != nil {
		return err
	}

	return nil
}
