package repository

import (
	"github.com/jmoiron/sqlx"
	m "github.com/venwex/threads/internal/models"
)

type PostRepository interface {
	ListsPosts() ([]m.Post, error)
	GetPost(id int) (m.Post, error)
	CreatePost(post m.Post) (m.Post, error)
	UpdatePost(id int, content string) (m.Post, error)
	DeletePost(id int) (m.Post, error)
}

type UserRepository interface {
	ListUsers()
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
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
