package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/venwex/threads/internal/handler/dto"
	m "github.com/venwex/threads/internal/models"
	"github.com/venwex/threads/internal/service"
	u "github.com/venwex/threads/internal/utils"
)

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{
		auth: svc,
	}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.SignUpRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error decoding sign up request json: %w, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.auth.SignUp(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		log.Printf("error signing up user %s: %v, path: %v", req.Username, err, r.URL.Path)
		if errors.Is(err, m.ErrUserAlreadyExists) {
			u.RenderError(w, http.StatusConflict, err.Error())
			return
		}

		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, user)
}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.SignInRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error decoding sign in request json: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.auth.SignIn(ctx, req.Login, req.Password)
	if err != nil {
		if errors.Is(err, m.ErrInvalidCredentials) {
			log.Printf("invalid credentials for user %s, path: %v", req.Login, r.URL.Path)
			u.RenderError(w, http.StatusBadRequest, err.Error())
			return
		}

		if errors.Is(err, m.ErrUserNotFound) {
			log.Printf("user %s not found, path: %v", req.Login, r.URL.Path)
			u.RenderError(w, http.StatusBadRequest, err.Error())
			return
		}

		log.Printf("error signing in user %s: %v, path: %v", req.Login, err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, tokens) // also return tokens
}

func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req dto.RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("error decoding refresh token request json: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	tokens, err := h.auth.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		if errors.Is(err, m.ErrInvalidRefreshToken) {
			u.RenderError(w, http.StatusUnauthorized, "invalid refresh token")
			return
		}

		if errors.Is(err, m.ErrUserNotFound) {
			u.RenderError(w, http.StatusUnauthorized, "invalid refresh token")
			return
		}

		log.Printf("error refreshing token: %v, path: %s", err, r.URL.Path)
		u.RenderError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	u.WriteJSON(w, http.StatusOK, tokens)
}
