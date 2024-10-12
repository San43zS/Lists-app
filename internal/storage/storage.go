package storage

import (
	"Lists-app/internal/broker/rabbit/api/notification"
	"Lists-app/internal/storage/api/user"
)

type Storage interface {
	User() user.User
	Notification() notification.Notification
}
