package repository

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	m "github.com/venwex/threads/internal/models"
)

type PostRepo struct {
	db *sqlx.DB
}

func NewPostRepo(db *sqlx.DB) *PostRepo {
	return &PostRepo{db: db}
}

func (repository *PostRepo) ListsPosts(ctx context.Context) ([]m.Post, error) {
	query := `
		select 
			p.post_id,
			p.author_id,
			u.username as author_username,
			p.content,
			p.created_at,
			p.updated_at,
			p.deleted_at
		from posts p
		join users u on u.user_id = p.author_id
		where p.deleted_at is null
		order by p.created_at desc;
	`

	var posts []m.Post
	if err := repository.db.SelectContext(ctx, &posts, query); err != nil {
		return nil, err
	}

	return posts, nil
}

func (repository *PostRepo) GetPost(ctx context.Context, postID uuid.UUID) (m.Post, error) {
	query := `
	select 
	    post_id, 
	    author_id, 
	    content, 
	    created_at, 
	    updated_at, 
	    deleted_at 
	from posts 
	where (post_id = $1 and deleted_at is null)
	`

	if postID == uuid.Nil {
		return m.Post{}, m.ErrInvalidID
	}

	var post m.Post
	if err := repository.db.GetContext(ctx, &post, query, postID); err != nil {
		return m.Post{}, err
	}

	return post, nil
}

func (repository *PostRepo) CreatePost(ctx context.Context, post m.Post) (m.Post, error) {
	query := `
		with created_post as (
			insert into posts (author_id, content)
			values ($1, $2)
			returning post_id, author_id, content, created_at, updated_at, deleted_at
		)
		select
			cp.post_id,
			cp.author_id,
			u.username as author_username,
			cp.content,
			cp.created_at,
			cp.updated_at,
			cp.deleted_at
		from created_post cp
		join users u on u.user_id = cp.author_id;
	`

	var createdPost m.Post
	err := repository.db.GetContext(ctx, &createdPost, query, post.AuthorID, post.Content)
	if err != nil {
		return m.Post{}, err
	}

	return createdPost, nil
}

func (repository *PostRepo) UpdatePost(ctx context.Context, postID uuid.UUID, content string) (m.Post, error) { // wrong
	if strings.TrimSpace(content) == "" {
		return m.Post{}, m.ErrBlankContent
	}

	if postID == uuid.Nil {
		return m.Post{}, m.ErrInvalidID
	}

	query := `update posts set content = $1, updated_at = now() where (post_id = $2 and deleted_at is null) returning post_id, author_id, content, created_at, updated_at, deleted_at`

	var post m.Post

	if err := repository.db.GetContext(ctx, &post, query, content, postID); err != nil {
		return m.Post{}, err
	}

	return post, nil
}

func (repository *PostRepo) DeletePost(ctx context.Context, postID uuid.UUID) (m.Post, error) {
	query := `update posts set deleted_at = now() where post_id = $1 returning post_id, author_id, content, created_at, updated_at, deleted_at;`

	if postID == uuid.Nil {
		return m.Post{}, m.ErrInvalidID
	}

	var post m.Post
	if err := repository.db.GetContext(ctx, &post, query, postID); err != nil {
		return m.Post{}, err
	}

	return post, nil
}
