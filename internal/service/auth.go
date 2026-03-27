package service

import "github.com/venwex/threads/internal/repository"

type AuthService struct {
	auth repository.UserRepository
}

func NewAuthService(auth repository.UserRepository) *AuthService {
	return &AuthService{auth: auth}
}
