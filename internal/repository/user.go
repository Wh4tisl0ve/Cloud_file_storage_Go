package repository

import (
	"database/sql"
	"fmt"

	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/entity"
)

type User struct {
	Conn *sql.DB
}

func (u *User) Save(usr *entity.User) (int, error) {
	query := "INSERT INTO users(username, password) VALUES($1, $2) RETURNING id"
	var id int

	row := u.Conn.QueryRow(query, usr.Username, usr.Password)
	if err := row.Scan(&id); err != nil {
		return 0, fmt.Errorf("Save user: %v", err)
	}

	return int(id), nil
}

func (u *User) Find(id int) (*entity.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = $1"
	row := u.Conn.QueryRow(query, id)
	usr := &entity.User{}

	if err := row.Scan(&usr.Id, &usr.Username, &usr.Password); err != nil {
		return nil, fmt.Errorf("Find user: %v", err)
	}

	return usr, nil
}
