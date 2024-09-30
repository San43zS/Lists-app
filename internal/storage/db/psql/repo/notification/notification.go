package notification

import (
	notification3 "Lists-app/internal/model/notification"
	notification2 "Lists-app/internal/storage/api/notification"
	"github.com/jmoiron/sqlx"
)

type notification struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) notification2.Notification {
	return notification{
		db: db,
	}
}

func (n notification) AddNotification(notification notification3.Notification) error {
	return nil
}

func (n notification) GetNotificationByUserId(Id int) (notification3.Notification, error) {
	return notification3.Notification{}, nil
}

func (n notification) GetAllNotificationsTTLon() ([]notification3.Notification, error) {
	return nil, nil
}

func (n notification) GetAllNotificationsTTLoff() ([]notification3.Notification, error) {
	return nil, nil
}

func (n notification) DeleteNotification(notification notification3.Notification) error {
	return nil
}
