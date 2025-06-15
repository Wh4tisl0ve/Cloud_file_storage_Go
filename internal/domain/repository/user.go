package user

import (
	"github.com/Wh4tisl0ve/Cloud_file_storage_Go/internal/domain/entity"
)

type Repository interface {
	Create(user entity.User) error
	User(id int) (*entity.User, error)
}
