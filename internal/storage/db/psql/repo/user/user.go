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
	var existingUser user22.User
	query := "SELECT * FROM users WHERE email = $1 AND password = $2 AND username = $3"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return false, err
	}

	err = stmt.QueryRow(user.Email, user.Password, user.Username).Scan(&existingUser.Id, &existingUser.Email, &existingUser.Username, &existingUser.Password)

	if err != nil || user.Email != existingUser.Email && user.Password != existingUser.Password && user.Username != existingUser.Username {
		return false, err
	}
	return true, nil
}

func (r repository) Insert(ctx context.Context, user user22.User) error {
	query := "INSERT INTO users (id, email, username, password) VALUES (DEFAULT, $1, $2, $3) RETURNING id"

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
