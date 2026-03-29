package handler

import "net/http"

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {}
