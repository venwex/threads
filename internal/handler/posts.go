package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	mw "github.com/venwex/threads/internal/middleware"
	"github.com/venwex/threads/internal/service"
	u "github.com/venwex/threads/internal/utils"
	ws "github.com/venwex/threads/internal/websockets"
)

type PostHandler struct {
	svc *service.PostService
	Hub *ws.Hub
}

func NewPostHandler(svc *service.PostService, hub *ws.Hub) *PostHandler {
	return &PostHandler{svc: svc, Hub: hub}
}

func (handler *PostHandler) ListPosts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	resp, err := handler.svc.ListPosts(ctx)
	if err != nil {
		log.Printf("error listing posts: %v\n, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, resp)
}

func (handler *PostHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	postID, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting post post id: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return

	}

	ctx := r.Context()
	post, err := handler.svc.GetPost(ctx, postID)
	if err != nil {
		log.Printf("error getting post: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, post)
}

func (handler *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	userID, ok := mw.GetUserID(r.Context())
	log.Printf("userID from context: %s, path: %s", userID.String(), r.URL.Path)
	if !ok {
		log.Println("error getting user id:, path: %v", r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, "error getting user id")
		return
	}

	post, err := u.DecodePost(r)
	if err != nil {
		log.Printf("error decoding post: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, fmt.Errorf("error during decoding (creating) post: %v", err).Error())
		return
	}

	post.AuthorID = userID

	ctx := r.Context()
	post, err = handler.svc.CreatePost(ctx, post)
	if err != nil {
		log.Printf("error creating post: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data, _ := json.Marshal(post)
	handler.Hub.Broadcast <- data

	u.WriteJSON(w, http.StatusCreated, post)
}

func (handler *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	postID, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting post id: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, fmt.Errorf("error during decoding (update) post: %v", err).Error())
		return
	}

	post, err := u.DecodePost(r)
	if err != nil {
		log.Printf("error decoding post: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	post, err = handler.svc.UpdatePost(ctx, postID, post.Content)
	if err != nil {
		log.Printf("error updating post: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, post)
}

func (handler *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	postID, err := u.GetID(r)
	if err != nil {
		log.Printf("error getting post id: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusBadRequest, fmt.Errorf("error during decoding (delete) post: %v", err).Error())
		return
	}

	ctx := r.Context()
	post, err := handler.svc.DeletePost(ctx, postID)
	if err != nil {
		log.Printf("error deleting post: %v, path: %v", err, r.URL.Path)
		u.RenderError(w, http.StatusNotFound, err.Error())
		return
	}

	u.WriteJSON(w, http.StatusOK, post)
}
