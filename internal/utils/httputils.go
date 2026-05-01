package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	m "github.com/venwex/threads/internal/models"
)

type H map[string]any

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("error during writing json:", err)
	}
}

func RenderError(w http.ResponseWriter, status int, text string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(H{"error": text}); err != nil {
		log.Println("error during json encoding: ", text)
	}
}

func GetID(r *http.Request) (uuid.UUID, error) {
	idStr := r.PathValue("id")

	id, err := uuid.Parse(idStr)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func DecodePost(r *http.Request) (m.Post, error) {
	var post m.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		return m.Post{}, err
	}

	return post, nil
}
