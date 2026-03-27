package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	e "github.com/venwex/threads/internal/errors"
	m "github.com/venwex/threads/internal/models"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (repository *PostRepo) ListsPosts() ([]m.Post, error) {
	var posts []m.Post

	if err := repository.db.Select(&posts, "select * from posts"); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *PostRepo) GetPost(id int) (m.Post, error) {
	if id <= 0 {
		return m.Post{}, e.ErrInvalidID
	}

	var post m.Post

	if err := repository.db.Get(&post, "select * from posts where id = $1", id); err != nil {
		return m.Post{}, err
	}

	return post, nil
}

func (repository *PostRepo) CreatePost(post m.Post) (m.Post, error) {
	err := repository.db.Get(&post, "insert into posts (author_id, content) values ($1, $2) returning id, author_id, content, created_at, updated_at", post.AuthorID, post.Content)
	if err != nil {
		return m.Post{}, err
	}

	return post, nil
}

func (repository *PostRepo) UpdatePost(id int, content string) (m.Post, error) { // wrong
	if len(content) == 0 {
		return m.Post{}, e.ErrBlankContent
	}

	var post m.Post

	if err := repository.db.Get(&post, "update posts set content = $1, updated_at = $2 where id = $3 returning id, author_id, content, created_at, updated_at", content, time.Now(), id); err != nil {
		return m.Post{}, err
	}

	return post, nil
}

func (repository *PostRepo) DeletePost(id int) (m.Post, error) {
	if id <= 0 {
		return m.Post{}, e.ErrInvalidID
	}

	var post m.Post

	if err := repository.db.Get(&post, "delete from posts where id = $1 returning id, author_id, content, created_at, updated_at", id); err != nil {
		return m.Post{}, err
	}

	return post, nil
}
