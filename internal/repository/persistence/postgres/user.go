package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/entity"
)

type Repository struct {
	Conn *sql.DB
}

func New(conn *sql.DB) *Repository {
	return &Repository{conn}
}

func (r Repository) Create(user *entity.User) error {
	fmt.Print("ura")
	return nil
}

func (r Repository) User(id int) (*entity.User, error) {
	fmt.Print("ura")
	return nil, nil
}
