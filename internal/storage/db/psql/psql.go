package psql

import (
	"Lists-app/internal/storage"
	"Lists-app/internal/storage/api/notification"
	user2 "Lists-app/internal/storage/api/user"
	"Lists-app/internal/storage/config"
	notification2 "Lists-app/internal/storage/db/psql/repo/notification"
	"Lists-app/internal/storage/db/psql/repo/user"
	"github.com/jmoiron/sqlx"
)

type Store struct {
	user         user2.User
	notification notification.Notification
}

func New(config config.Config) (storage.Storage, error) {
	db, err := sqlx.Open(config.Driver, config.URL)

	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	//!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	return &Store{
		user:         user.New(db),
		notification: notification2.New(db),
	}, nil
}

func (s *Store) User() user2.User {
	return s.user
}

func (s *Store) Notification() notification.Notification {
	return s.notification
}