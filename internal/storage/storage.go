package storage

import (
	"Lists-app/internal/model/user"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	user2 "os/user"
)

type Storage struct {
	db *sqlx.DB
}

type DB interface {
	VerifyUser(user user.User) error
	InsertUser(user user.User) error
}

type dbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getDBConfig() dbConfig {
	return dbConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_NAME"),
	}
}

// есть ли такой user уже
func (s *Storage) VerifyUser(user user.User) error {
	// Выполняем запрос к базе данных
	query := "SELECT * FROM users WHERE email = $1"
	var existingUser user2.User
	err := s.db.Get(&existingUser, query, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return err
	}
	// Если пользователь найден, возвращаем nil
	return nil
}

// добавляем нового пользователя
func (s *Storage) InsertUser(user user.User) error {
	query := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)"

	_, err := s.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
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
