package user

import (
	user22 "Lists-app/internal/model/user"
	user2 "Lists-app/internal/storage/api/user"
	"context"
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) user2.User {
	return repository{
		db: db,
	}
}

func (r repository) GetById(ctx context.Context, Id int) (user22.User, error) {
	query := "SELECT * FROM users WHERE Id = $1"
	var existingUser user22.User

	err := r.db.Get(&existingUser, query, Id)
	if err != nil {
		return user22.User{}, err
	}
	return existingUser, nil
}

func (r repository) Verify(ctx context.Context, user user22.User) (bool, error) {
	// Выполняем запрос к базе данных
	query := "SELECT * FROM users WHERE email = $1 AND password = $2 AND username = $3"
	var existingUser user2.User
	err := r.db.Get(&existingUser, query, user.Email, user.Password, user.Username)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r repository) Insert(ctx context.Context, user user22.User) error {
	query := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)"

	_, err := r.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r repository) Delete(ctx context.Context, user user22.User) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, user.Id)
	if err != nil {
		return err
	}
	return nil
}
