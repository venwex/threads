package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	m "github.com/venwex/threads/internal/models"
)

type PostRepository interface {
	ListsPosts(ctx context.Context) ([]m.Post, error)
	GetPost(ctx context.Context, id int) (m.Post, error)
	CreatePost(ctx context.Context, post m.Post) (m.Post, error)
	UpdatePost(ctx context.Context, id int, content string) (m.Post, error)
	DeletePost(ctx context.Context, id int) (m.Post, error)
}

type UserRepository interface {
	ListUsers()
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type AuthService interface {
	SignUp(context.Context) (m.User, error)
	SignIn(ctx context.Context, username string, password string) (m.User, error)
}

type Repository struct {
	Post PostRepository
	User UserRepository
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Post: NewPostRepo(db),
		User: NewUserRepo(db),
	}
}
