package notification

import "Lists-app/internal/model/user"

type Notification interface {
	GetNotification(user user.User) error
}
