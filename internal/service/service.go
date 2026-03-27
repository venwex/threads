package service

import "github.com/venwex/threads/internal/repository"

type Service struct {
	Posts *PostService
	Users *UserService
	Auth  *AuthService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Posts: NewPostService(repos.Post),
		Users: NewUserService(repos.User),
		Auth:  NewAuthService(repos.User),
	}
}
