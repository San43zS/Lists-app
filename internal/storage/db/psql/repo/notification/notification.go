package notification

import (
	notification3 "Lists-app/internal/model/notification"
	notification2 "Lists-app/internal/storage/api/notification"
)

type notification struct {
	notification2.Notification
}

func (n *notification) AddNotification(notification notification3.Notification) error {

}

func (n *notification) GetNotificationByUserId(Id int) (notification3.Notification, error) {

}

func (n *notification) GetAllNotificationsTTLon() ([]notification3.Notification, error) {

}

func (n *notification) GetAllNotificationsTTLoff() ([]notification3.Notification, error) {

}
