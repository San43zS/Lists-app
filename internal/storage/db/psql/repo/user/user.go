package user

import (
	user22 "Lists-app/internal/model/user"
	user2 "Lists-app/internal/storage/api/user"
	"github.com/jmoiron/sqlx"
)

type user struct {
	db   *sqlx.DB
	user user2.User
}

func (u *user) GetById(Id int) (user22.User, error) {
	query := "SELECT * FROM users WHERE Id = $1"
	var existingUser user22.User

	err := u.db.Get(&existingUser, query, Id)
	if err != nil {
		return user22.User{}, err
	}
	return existingUser, nil
}

func (u *user) Verify(user user22.User) (bool, error) {
	// Выполняем запрос к базе данных
	query := "SELECT * FROM users WHERE email = $1 AND password = $2 AND username = $3"
	var existingUser user2.User
	err := u.db.Get(&existingUser, query, user.Email, user.Password, user.Username)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *user) Insert(user user22.User) error {
	query := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)"

	_, err := u.db.Exec(query, user.Email, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}
