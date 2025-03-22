package storage

import (
	"notify-service/internal/storage/api/user"
)

type Storage interface {
	User() user.User
}
