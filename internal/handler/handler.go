package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/venwex/threads/internal/service"
	u "github.com/venwex/threads/internal/utils"
)

type Handler struct {
	svc *service.Service
}

type H map[string]any

func NewHandler(svc *service.Service) *Handler {
	return &Handler{svc: svc}
}

func (handler *Handler) Health(w http.ResponseWriter, r *http.Request) {
	u.WriteJSON(w, http.StatusOK, H{
		"healthy": true,
	})
}

func (handler *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {

}

func (handler *Handler) ListPosts(w http.ResponseWriter, r *http.Request) {
	resp, err := handler.svc.ListPosts()
	if err != nil {
		log.Println("error listing posts: ", err)
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, resp)
}

func (handler *Handler) GetPost(w http.ResponseWriter, r *http.Request) {
	id, err := u.GetID(r)
	if err != nil {
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err := handler.svc.GetPost(id)
	if err != nil {
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, post)
}

func (handler *Handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	post, err := u.DecodePost(r)
	if err != nil {
		u.RenderError(w, http.StatusBadRequest, fmt.Errorf("error during decoding (creating) post: %v", err).Error())
		return
	}

	post, err = handler.svc.CreatePost(post)
	if err != nil {
		u.RenderError(w, http.StatusInternalServerError, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusCreated, post)
}

func (handler *Handler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	id, err := u.GetID(r)
	if err != nil {
		u.RenderError(w, http.StatusBadRequest, fmt.Errorf("error during decoding (update) post: %v", err).Error())
		return
	}

	post, err := u.DecodePost(r)
	if err != nil {
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	post, err = handler.svc.UpdatePost(id, post.Content)
	if err != nil {
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, post)
}

func (handler *Handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	id, err := u.GetID(r)
	if err != nil {
		u.RenderError(w, http.StatusBadRequest, fmt.Errorf("error during decoding (delete) post: %v", err).Error())
		return
	}

	post, err := handler.svc.DeletePost(id)
	if err != nil {
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, post)
}

func (handler *Handler) SignIn(w http.ResponseWriter, r *http.Request) {}

func (handler *Handler) SignUp(w http.ResponseWriter, r *http.Request) {}
