package notification

import "Lists-app/internal/model/notification"

type Notification interface {
	GetNotificationByUserId(Id int) (notification.Notification, error)
	GetAllNotificationsTTLoff() ([]notification.Notification, error)
	AddNotification(notification Notification) error
	GetAllNotificationsTTLon() ([]notification.Notification, error)
	DeleteNotification(notification notification.Notification) error
}
