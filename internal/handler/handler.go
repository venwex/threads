package handler

import (
	"net/http"

	"github.com/venwex/threads/internal/service"
	u "github.com/venwex/threads/internal/utils"
)

type Handler struct {
	Posts *PostHandler
	Users *UserHandler
	Auth  *AuthoHandler
}

type H map[string]any

func NewHandler(svc *service.Service) *Handler {
	return &Handler{
		Posts: NewPostHandler(svc.Posts),
		Users: NewUsersHandler(svc.Users),
	}
}

func (handler *Handler) Health(w http.ResponseWriter, r *http.Request) {
	u.WriteJSON(w, http.StatusOK, H{
		"healthy": true,
	})
}
