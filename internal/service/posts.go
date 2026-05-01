package service

import (
	"context"

	"github.com/google/uuid"
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

func (service *PostService) GetPost(ctx context.Context, postID uuid.UUID) (m.Post, error) {
	return service.posts.GetPost(ctx, postID)
}

func (service *PostService) CreatePost(ctx context.Context, post m.Post) (m.Post, error) {
	return service.posts.CreatePost(ctx, post)
}

func (service *PostService) UpdatePost(ctx context.Context, postID uuid.UUID, content string) (m.Post, error) {
	return service.posts.UpdatePost(ctx, postID, content)
}

func (service *PostService) DeletePost(ctx context.Context, postID uuid.UUID) (m.Post, error) {
	return service.posts.DeletePost(ctx, postID)
}
