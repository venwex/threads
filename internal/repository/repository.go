package repository

import (
	"github.com/jmoiron/sqlx"
)

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
