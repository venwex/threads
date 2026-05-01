package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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

func (s *AuthRepo) SignUp(ctx context.Context, username, passwordHash, email string) (m.User, error) {
	query := `
		insert into users (username, password_hash, email) 
		values ($1, $2, $3)
		returning user_id, username, email, role, created_at, updated_at, updated_at;
	`

	var user m.User
	err := s.db.GetContext(ctx, &user, query, username, passwordHash, email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthRepo) GetUserByLogin(ctx context.Context, login string) (m.User, error) {
	query := `
		select user_id, username, email, password_hash, role, created_at, updated_at from users where (username = $1 or email = $1) and deleted_at is null;
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

func (s *AuthRepo) FindRefreshToken(ctx context.Context, refreshHash string) (uuid.UUID, error) { // userID, error
	query := `
		select user_id from refresh_tokens where (refresh_hash = $1 and revoked_at is null and expires_at > now());
	`

	var userID uuid.UUID
	err := s.db.GetContext(ctx, &userID, query, refreshHash)
	if err != nil {
		return uuid.Nil, m.ErrInvalidRefreshToken
	}

	return userID, nil
}

func (s *AuthRepo) GetUserByID(ctx context.Context, userID uuid.UUID) (*m.User, error) {
	query := `
		select user_id, username, email, password_hash, role, created_at, updated_at from users where (user_id = $1 and deleted_at is null);
	`

	var user m.User
	err := s.db.GetContext(ctx, &user, query, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, m.ErrUserNotFound
		}

		return nil, err
	}

	return &user, nil
}

func (s *AuthRepo) RotateRefreshToken(
	ctx context.Context,
	userID uuid.UUID,
	oldRefreshHash,
	newRefreshHash string,
	expiresAt time.Time) error {

	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}

	defer tx.Rollback()

	revokeQuery := `
		update refresh_tokens 
		set revoked_at = now() 
		where (user_id = $1 
		   and refresh_hash = $2 
		   and revoked_at is null 
		   and expires_at > now());
	`

	res, err := tx.ExecContext(ctx, revokeQuery, userID, oldRefreshHash)
	if err != nil {
		return fmt.Errorf("revoke old refresh token: %w", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected by revoke old refresh token: %w", err)
	}
	if rowsAffected == 0 {
		return m.ErrInvalidRefreshToken
	}

	insertQuery := `
		insert into refresh_tokens(user_id, refresh_hash, expires_at)
		values ($1, $2, $3);
	`

	_, err = tx.ExecContext(ctx, insertQuery, userID, newRefreshHash, expiresAt)
	if err != nil {
		return fmt.Errorf("insert new refresh token: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit refresh token rotation tx: %w", err)
	}

	return nil
}
