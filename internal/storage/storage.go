package storage

import (
	"Lists-app/internal/storage/api/user"
)

type Storage interface {
	User() user.User
}
