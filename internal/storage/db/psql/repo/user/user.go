package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	error2 "notify-service/internal/model/error"
	user22 "notify-service/internal/model/user"
	user2 "notify-service/internal/storage/api/user"
	"notify-service/internal/storage/config"
	"notify-service/pkg/encrypt"
)

type repository struct {
	db      *sqlx.DB
	encrypt encrypt.Service
}

func New(db *sqlx.DB, encrypt encrypt.Service) user2.User {
	return repository{
		db:      db,
		encrypt: encrypt,
	}
}

func (r repository) GetById(ctx context.Context, Id int) (user22.User, error) {
	var existingUser user22.User

	query := "SELECT * FROM users WHERE Id = $1"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return user22.User{}, fmt.Errorf("failed to prepare query: %w", err)
	}

	err = stmt.QueryRowContext(ctx, Id).Scan(
		&existingUser.Id,
		&existingUser.Email,
		&existingUser.Username,
		&existingUser.Password,
	)

	if err != nil {
		return user22.User{}, fmt.Errorf("failed to query row: %w", err)
	}

	return existingUser, nil
}

func (r repository) SignIn(ctx context.Context, user user22.User) error {
	var existingUser user22.User

	query := `SELECT * FROM users WHERE email = $1 AND password = $2`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	err = stmt.QueryRowContext(
		ctx,
		user.Email,
		user.Password,
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

	if user.Email != existingUser.Email {
		return error2.ErrWrongEmail
	}

	if r.encrypt.Password(user.Password) != existingUser.Password {
		return error2.ErrWrongPassword
	}

	return nil
}

func (r repository) SignUp(ctx context.Context, user user22.User) error {
	user.Password = r.encrypt.Password(user.Password)
	query := "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)"

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, user.Email, user.Username, user.Password)
	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && string(pgErr.Code) == config.GetUniqueViolationErr() {
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
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	_, err = stmt.ExecContext(ctx, user.Id)
	if err != nil {
		return error2.ErrUnknown
	}

	return nil
}
