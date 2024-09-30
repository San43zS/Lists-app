package notification

type Notification interface {
	GetNotificationByUserId(Id int) (Notification, error)
	GetAllNotificationsTTLoff() ([]Notification, error)
	AddNotification(Notification Notification) error
	GetAllNotificationsTTLon() ([]Notification, error)
	DeleteNotification(Notification Notification) error
}
