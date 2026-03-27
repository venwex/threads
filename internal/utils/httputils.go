package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

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

func GetID(r *http.Request) (int, error) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
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
