package handler

import "net/http"

type AuthoHandler struct {
}

func NewAuthoHandler() *AuthoHandler {
	return &AuthoHandler{}
}

func (h *AuthoHandler) SignUp(w http.ResponseWriter, r *http.Request) {}

func (h *AuthoHandler) SignIn(w http.ResponseWriter, r *http.Request) {}
