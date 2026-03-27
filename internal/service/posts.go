package service

import (
	m "github.com/venwex/threads/internal/models"
	"github.com/venwex/threads/internal/repository"
)

type PostService struct {
	posts repository.PostRepository
}

func NewPostService(posts repository.PostRepository) *PostService {
	return &PostService{posts: posts}
}

func (service *PostService) ListPosts() ([]m.Post, error) {
	return service.posts.ListsPosts()
}

func (service *PostService) GetPost(id int) (m.Post, error) {
	return service.posts.GetPost(id)
}

func (service *PostService) CreatePost(post m.Post) (m.Post, error) {
	return service.posts.CreatePost(post)
}

func (service *PostService) UpdatePost(id int, content string) (m.Post, error) {
	return service.posts.UpdatePost(id, content)
}

func (service *PostService) DeletePost(id int) (m.Post, error) {
	return service.posts.DeletePost(id)
}
