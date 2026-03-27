package handler

import (
	"net/http"

	"github.com/venwex/threads/internal/service"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUsersHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc: svc}
}

func (handler *UserHandler) ListUsers(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}
