package user

import (
	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/domain/repository"
)

type Login struct {
	r user.Repository
	// токен сервис
}

func New(r user.Repository) *Login {
	return &Login{
		r: r,
	}
}

func (uc *Login) Execute(username string, password string) (token string, err error) {
	uc.r.User(1)
	return "tok", nil
}
