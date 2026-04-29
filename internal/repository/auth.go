package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	m "github.com/venwex/threads/internal/models"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (s *AuthRepo) ExistsByUsernameOrEmail(ctx context.Context, username, email string) (bool, error) {
	query := `
		select exists(select 1 from users where (username = $1 or email = $2) and deleted_at is null);
	`

	var exists bool
	err := s.db.GetContext(ctx, &exists, query, username, email)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func (s *AuthRepo) SignUp(ctx context.Context, username, password, email string) (m.User, error) {
	query := `
		insert into users (username, password_hash, email) 
		values ($1, $2, $3)
		returning id, username, email;
	` // изменить in db threads_db from password_hash to password

	var user m.User
	err := s.db.GetContext(ctx, &user, query, username, password, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthRepo) GetUser(ctx context.Context, login string) (m.User, error) {
	query := `
		select id, username, email, password_hash from users where (username = $1 or email = $1);
	`

	var user m.User
	err := s.db.GetContext(ctx, &user, query, login)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return m.User{}, m.ErrUserNotFound
		}

		return user, err
	}

	return user, nil
}

func (s *AuthRepo) SaveRefreshToken(ctx context.Context, userID uuid.UUID, refreshHash string, expiresAt time.Time) error {
	query := `insert into refresh_tokens(user_id, refresh_hash, expires_at) values ($1, $2, $3);`
	_, err := s.db.ExecContext(ctx, query, userID, refreshHash, expiresAt)
	if err != nil {
		return err
	}

	return nil
}
