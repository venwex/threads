package handler

import (
	"net/http"

	"github.com/venwex/threads/internal/service"
	u "github.com/venwex/threads/internal/utils"
	ws "github.com/venwex/threads/internal/websockets"
)

type Handler struct {
	Posts *PostHandler
	Users *UserHandler
	Auth  *AuthHandler
}

type H map[string]any

func NewHandler(svc *service.Service, hub *ws.Hub) *Handler {
	return &Handler{
		Posts: NewPostHandler(svc.Posts, hub),
		Users: NewUsersHandler(svc.Users),
		Auth:  NewAuthHandler(svc.Auth),
	}
}

func (handler *Handler) Health(w http.ResponseWriter, r *http.Request) {
	u.WriteJSON(w, http.StatusOK, H{
		"healthy": true,
	})
}
