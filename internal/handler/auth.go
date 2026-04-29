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
		log.Println("error decoding sign up request json:", err)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.auth.SignUp(ctx, req.Username, req.Password, req.Email)
	if err != nil {
		log.Printf("error signing up user %s: %v", req.Username, err)
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
		log.Println("error decoding sign in request json:", err)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.auth.SignIn(ctx, req.Login, req.Password)
	if err != nil {
		if errors.Is(err, m.ErrInvalidCredentials) {
			log.Printf("invalid credentials for user %s", req.Login)
			u.RenderError(w, http.StatusBadRequest, err.Error())
			return
		}

		log.Printf("error signing in user %s: %v", req.Login, err)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, user) // also add tokens
}
