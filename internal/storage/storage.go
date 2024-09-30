package storage

import (
	"Lists-app/internal/storage/api/notification"
	user22 "Lists-app/internal/storage/api/user"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Storage interface {
	User() user22.User
	Notification() notification.Notification
}

func New() DB {
	dbEssence := getDBConfig()

	// Создаем строку подключения к базе данных
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbEssence.Host, dbEssence.Port, dbEssence.User, dbEssence.Password, dbEssence.Database)

	// Подключаемся к базе данных
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil
	}
	defer db.Close()

	return &Storage{
		db: db,
	}
}
