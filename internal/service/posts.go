package service

import (
	"context"

	m "github.com/venwex/threads/internal/models"
	"github.com/venwex/threads/internal/repository"
)

type PostService struct {
	posts repository.PostRepository
}

func NewPostService(posts repository.PostRepository) *PostService {
	return &PostService{posts: posts}
}

func (service *PostService) ListPosts(ctx context.Context) ([]m.Post, error) {
	return service.posts.ListsPosts(ctx)
}

func (service *PostService) GetPost(ctx context.Context, id int) (m.Post, error) {
	return service.posts.GetPost(ctx, id)
}

func (service *PostService) CreatePost(ctx context.Context, post m.Post) (m.Post, error) {
	return service.posts.CreatePost(ctx, post)
}

func (service *PostService) UpdatePost(ctx context.Context, id int, content string) (m.Post, error) {
	return service.posts.UpdatePost(ctx, id, content)
}

func (service *PostService) DeletePost(ctx context.Context, id int) (m.Post, error) {
	return service.posts.DeletePost(ctx, id)
}
