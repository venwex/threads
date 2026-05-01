package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID  `json:"user_id" db:"user_id"`
	Username  string     `json:"username" db:"username"`
	Email     string     `json:"email" db:"email"`
	Password  string     `json:"-" db:"password_hash"`
	Role      string     `json:"role" db:"role"`
	CreatedAt time.Time  `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time  `json:"updatedAt" db:"updated_at"`
	DeletedAt *time.Time `json:"deletedAt" db:"deleted_at"`
}

type Post struct {
	PostID         uuid.UUID  `json:"post_id" db:"post_id"`
	AuthorID       uuid.UUID  `json:"author_id" db:"author_id"`
	AuthorUsername string     `json:"author_username" db:"author_username"`
	Content        string     `json:"content" db:"content"`
	CreatedAt      time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at" db:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at" db:"deleted_at"`
}
