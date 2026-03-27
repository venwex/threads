package service

import (
	m "github.com/venwex/threads/internal/models"
	"github.com/venwex/threads/internal/repository"
)

type Service struct {
	repo *repository.UserRepo
}

func NewService(repo *repository.UserRepo) *Service {
	return &Service{repo: repo}
}

func (service *Service) ListUsers() {}

func (service *Service) GetUser() {}

func (service *Service) CreateUser() {}

func (service *Service) UpdateUser() {}

func (service *Service) DeleteUser() {}

func (service *Service) ListPosts() ([]m.Post, error) {
	return service.repo.ListsPosts()
}

func (service *Service) GetPost(id int) (m.Post, error) {
	return service.repo.GetPost(id)
}

func (service *Service) CreatePost(post m.Post) (m.Post, error) {
	return service.repo.CreatePost(post)
}

func (service *Service) UpdatePost(id int, content string) (m.Post, error) {
	return service.repo.UpdatePost(id, content)
}

func (service *Service) DeletePost(id int) (m.Post, error) {
	return service.repo.DeletePost(id)
}
