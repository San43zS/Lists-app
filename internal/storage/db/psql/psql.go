package psql

import (
	"Lists-app/internal/storage"
	user2 "Lists-app/internal/storage/api/user"
	"Lists-app/internal/storage/config"
	"Lists-app/internal/storage/db/psql/repo/user"
	"Lists-app/pkg/encrypt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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
