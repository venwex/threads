package service

import (
	"github.com/venwex/threads/internal/auth"
	"github.com/venwex/threads/internal/repository"
)

type Service struct {
	Posts *PostService
	Users *UserService
	Auth  *AuthService
}

func NewService(repos *repository.Repository, tokenManager *auth.TokenManager) *Service {
	return &Service{
		Posts: NewPostService(repos.Post),
		Users: NewUserService(repos.User),
		Auth:  NewAuthService(repos.Auth, tokenManager),
	}
}
