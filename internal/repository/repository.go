package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	m "github.com/venwex/threads/internal/models"
)

type PostRepository interface {
	ListsPosts(ctx context.Context) ([]m.Post, error)
	GetPost(ctx context.Context, id uuid.UUID) (m.Post, error)
	CreatePost(ctx context.Context, post m.Post) (m.Post, error)
	UpdatePost(ctx context.Context, id uuid.UUID, content string) (m.Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) (m.Post, error)
}

type UserRepository interface {
	ListUsers()
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type AuthRepository interface {
	SignUp(ctx context.Context, username, password, email string) (m.User, error)
	SaveRefreshToken(ctx context.Context, userID uuid.UUID, refreshHash string, expiresAt time.Time) error
	GetUserByLogin(ctx context.Context, login string) (m.User, error)
	ExistsByUsernameOrEmail(ctx context.Context, username, email string) (bool, error)
	FindRefreshToken(ctx context.Context, refreshHash string) (uuid.UUID, error)
	GetUserByID(ctx context.Context, userID uuid.UUID) (*m.User, error)
	RotateRefreshToken(ctx context.Context, userID uuid.UUID, oldRefreshHash, newRefreshHash string, expiresAt time.Time) error
}

type Repository struct {
	Post PostRepository
	User UserRepository
	Auth AuthRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: NewPostRepo(db),
		User: NewUserRepo(db),
		Auth: NewAuthRepo(db),
	}
}
