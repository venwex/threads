package repository

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	ListUsers()
	GetUser()
	CreateUser()
	UpdateUser()
	DeleteUser()
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (repository *UserRepo) ListUsers() {}

func (repository *UserRepo) GetUser() {}

func (repository *UserRepo) CreateUser() {}

func (repository *UserRepo) UpdateUser() {}

func (repository *UserRepo) DeleteUser() {}
