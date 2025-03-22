package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"notify-service/internal/storage"
	user2 "notify-service/internal/storage/api/user"
	"notify-service/internal/storage/config"
	"notify-service/internal/storage/db/psql/repo/user"
	"notify-service/pkg/encrypt"
)

type Store struct {
	user user2.User
}

func New(config *config.Config) (storage.Storage, error) {
	db, err := sqlx.Connect(config.Driver, config.URL)
	if err != nil {
		return nil, err
	}

	return &Store{
		user: user.New(db, encrypt.New()),
	}, nil
}

func (s Store) User() user2.User {
	return s.user
}
