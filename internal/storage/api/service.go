package api

import (
	"Lists-app/internal/storage/api/notification"
	"Lists-app/internal/storage/api/user"
)

type Service interface {
	notification.Notification
	user.User
}
