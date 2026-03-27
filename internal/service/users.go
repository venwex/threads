package service

import "github.com/venwex/threads/internal/repository"

type UserService struct {
	users repository.UserRepository
}

func NewUserService(users repository.UserRepository) *UserService {
	return &UserService{users: users}
}

func (service *PostService) ListUsers() {}

func (service *PostService) GetUser() {}

func (service *PostService) CreateUser() {}

func (service *PostService) UpdateUser() {}

func (service *PostService) DeleteUser() {}
