package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/venwex/threads/internal/handler"
	"github.com/venwex/threads/internal/repository"
	"github.com/venwex/threads/internal/service"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal("error init db: ", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("error during closing database connection: %v", err.Error())
			return
		}
	}(db)

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	mux := initRoutes(h)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

func initRoutes(h *handler.Handler) *http.ServeMux { // default crud
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", h.Health)

	mux.HandleFunc("POST /sign-up", h.SignUp)
	mux.HandleFunc("POST /sign-in", h.SignIn)

	mux.HandleFunc("GET /users", h.ListUsers)
	mux.HandleFunc("GET /user/{id}", h.GetUser)
	mux.HandleFunc("POST /user", h.CreateUser)
	mux.HandleFunc("PATCH /user/{id}", h.UpdateUser)
	mux.HandleFunc("DELETE /user/{id}", h.DeleteUser)

	mux.HandleFunc("GET /posts", h.ListPosts)
	mux.HandleFunc("GET /posts/{id}", h.GetPost)
	mux.HandleFunc("POST /posts", h.CreatePost)
	mux.HandleFunc("PATCH /posts/{id}", h.UpdatePost)
	mux.HandleFunc("DELETE /posts/{id}", h.DeletePost)

	return mux
}

func initDB() (*sqlx.DB, error) {
	dsn := "postgres://postgres:1234@localhost:5431/threads_db?sslmode=disable"

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to postgres")

	return db, nil
}
