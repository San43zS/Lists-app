package user

import (
	error2 "Lists-app/internal/model/error"
	user22 "Lists-app/internal/model/user"
	user2 "Lists-app/internal/storage/api/user"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
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
	var existingUser user22.User

	query := "SELECT * FROM users WHERE Id = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user22.User{}, err
	}

	err = stmt.QueryRowContext(ctx, Id).Scan(
		&existingUser.Id,
		&existingUser.Email,
		&existingUser.Username,
		&existingUser.Password,
	)

	if err != nil {
		return user22.User{}, err
	}

	return existingUser, nil
}

func (r repository) Verify(ctx context.Context, user user22.User) error {
	var existingUser user22.User

	query := `SELECT * FROM users WHERE email = $1 AND password = $2 AND username = $3`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	err = stmt.QueryRowContext(
		ctx,
		user.Email,
		user.Password,
		user.Username,
	).Scan(
		&existingUser.Id,
		&existingUser.Email,
		&existingUser.Username,
		&existingUser.Password,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return error2.ErrUserNotFound
		}

		return error2.ErrUnknown
	}

	if user.Email != existingUser.Email || user.Username != existingUser.Username || user.Password != existingUser.Password {
		return error2.ErrVerifyUser
	}

	return nil
}

func (r repository) Insert(ctx context.Context, user user22.User) error {
	query := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, query, user.Email, user.Username, user.Password)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == error2.UniqueViolationErr {
			return error2.UserAlreadyExistsErr
		}

		return error2.ErrUnknown
	}

	return nil
}

func (r repository) Delete(ctx context.Context, user user22.User) error {
	query := "DELETE FROM users WHERE id = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, query, user.Id)
	if err != nil {
		return error2.ErrUnknown
	}

	return nil
}
