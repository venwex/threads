package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/venwex/threads/internal/auth"
	m "github.com/venwex/threads/internal/models"
	"github.com/venwex/threads/internal/repository"
)

type AuthService struct {
	authRepo     repository.AuthRepo
	tokenManager *auth.TokenManager
}

func NewAuthService(auth repository.UserRepository, tokenManager *auth.TokenManager) *AuthService {
	return &AuthService{
		authRepo:     auth,
		tokenManager: tokenManager,
	}
}

func (s *AuthService) SignUp(ctx context.Context, username, password, email string) (m.User, error) {
	exists, err := s.authRepo.ExistsByUsernameOrEmail(ctx, username, email)
	if err != nil {
		return m.User{}, fmt.Errorf("error checking if user exists during sign-up: %w", err)
	}

	if exists {
		return m.User{}, m.ErrUserAlreadyExists
	}

	passwordHash, err := auth.HashPassword(password)
	if err != nil {
		return m.User{}, fmt.Errorf("error hashing password: %w", err)
	}

	user, err := s.authRepo.SignUp(ctx, username, passwordHash, email)
	if err != nil {
		return m.User{}, fmt.Errorf("error signing up user %s: %w", username, err)
	}

	return user, nil
}

func (s *AuthService) SignIn(ctx context.Context, login, password string) (m.AuthTokens, error) {
	user, err := s.authRepo.GetUser(ctx, login)
	if err != nil {
		if errors.Is(err, m.ErrInvalidCredentials) {
			return m.AuthTokens{}, m.ErrInvalidCredentials
		}
	}

	if !auth.CheckPasswordHash(password, user.Password) { // user.Password is its password_hash
		return m.AuthTokens{}, m.ErrInvalidCredentials
	}

	accessToken, err := s.tokenManager.GenerateAccessToken(user.ID, user.Username, user.Email, 30*time.Minute)
	if err != nil {
		return m.AuthTokens{}, fmt.Errorf("error generating access token for user %s: %w", user.Username, err)
	}

	refreshToken, refreshHash, err := auth.GenerateRefreshToken()
	if err != nil {
		return m.AuthTokens{}, fmt.Errorf("error generating refresh token for user %s: %w", user.Username, err)
	}

	err = s.authRepo.SaveRefreshToken(
		ctx,
		user.ID,
		refreshHash,
		time.Now().Add(30*24*time.Hour),
	)
	if err != nil {
		return m.AuthTokens{}, fmt.Errorf("save refresh token: %w", err)
	}

	return m.AuthTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
