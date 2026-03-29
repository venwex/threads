package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/venwex/threads/internal/handler"
	mw "github.com/venwex/threads/internal/middleware"
	"github.com/venwex/threads/internal/repository"
	"github.com/venwex/threads/internal/service"
	ws "github.com/venwex/threads/internal/websockets"
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

	hub := ws.NewHub()
	go hub.Run()

	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc, hub)

	mux := initRoutes(h)
	handlers := mw.Logging(mw.Cors(mux))
	
	log.Fatal(http.ListenAndServe(":8080", handlers))
}

func initRoutes(h *handler.Handler) *http.ServeMux { // default crud
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./index.html")
	})

	mux.HandleFunc("GET /health", h.Health)

	mux.HandleFunc("POST /sign-up", h.Auth.SignUp)
	mux.HandleFunc("POST /sign-in", h.Auth.SignIn)

	mux.HandleFunc("GET /users", h.Users.ListUsers)
	mux.HandleFunc("GET /user/{id}", h.Users.GetUser)
	mux.HandleFunc("POST /user", h.Users.CreateUser)
	mux.HandleFunc("PATCH /user/{id}", h.Users.UpdateUser)
	mux.HandleFunc("DELETE /user/{id}", h.Users.DeleteUser)

	mux.HandleFunc("GET /posts", h.Posts.ListPosts)
	mux.HandleFunc("GET /posts/{id}", h.Posts.GetPost)
	mux.HandleFunc("POST /posts", h.Posts.CreatePost)
	mux.HandleFunc("PATCH /posts/{id}", h.Posts.UpdatePost)
	mux.HandleFunc("DELETE /posts/{id}", h.Posts.DeletePost)

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(h.Posts.Hub, w, r)
	})

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
