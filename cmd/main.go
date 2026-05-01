package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/venwex/threads/internal/auth"
	"github.com/venwex/threads/internal/handler"
	mw "github.com/venwex/threads/internal/middleware"
	m "github.com/venwex/threads/internal/models"
	"github.com/venwex/threads/internal/repository"
	"github.com/venwex/threads/internal/service"
	ws "github.com/venwex/threads/internal/websockets"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found")
	}

	cfg := initCfg() // config for postgres

	db, err := initDB(cfg)
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

	tokenManger := auth.NewTokenManager(os.Getenv("JWT_TOKEN_SECRET"), os.Getenv("ISSUER"))

	repo := repository.NewRepository(db)
	svc := service.NewService(repo, tokenManger)
	h := handler.NewHandler(svc, hub)

	mux := initRoutes(h, tokenManger)
	handlers := mw.Logging(mw.Cors(mux))

	log.Fatal(http.ListenAndServe(":8080", handlers))
}

func initRoutes(h *handler.Handler, tokenManager *auth.TokenManager) *http.ServeMux {
	mux := http.NewServeMux()

	authMW := mw.AuthMiddleware(tokenManager)

	protected := func(pattern string, fn http.HandlerFunc) {
		mux.Handle(pattern, authMW(fn))
	}

	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	mux.HandleFunc("GET /health", h.Health)

	mux.HandleFunc("POST /sign-up", h.Auth.SignUp)
	mux.HandleFunc("POST /sign-in", h.Auth.SignIn)
	mux.HandleFunc("POST /refresh", h.Auth.RefreshToken)

	protected("GET /users", h.Users.ListUsers)
	protected("GET /users/{id}", h.Users.GetUser)
	protected("POST /users", h.Users.CreateUser)
	protected("PATCH /users/{id}", h.Users.UpdateUser)
	protected("DELETE /users/{id}", h.Users.DeleteUser)

	protected("GET /posts", h.Posts.ListPosts)
	protected("GET /posts/{id}", h.Posts.GetPost)
	protected("POST /posts", h.Posts.CreatePost)
	protected("PATCH /posts/{id}", h.Posts.UpdatePost)
	protected("DELETE /posts/{id}", h.Posts.DeletePost)

	protected("GET /ws", func(w http.ResponseWriter, r *http.Request) {
		ws.ServeWS(h.Posts.Hub, w, r)
	})

	return mux
}

func initDB(cfg m.PostgresConfig) (*sqlx.DB, error) {
	dsn := cfg.DSN()

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

func initCfg() m.PostgresConfig {
	return m.PostgresConfig{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSL_MODE"),
	}
}

//{
//"access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiOTU3YmM1OWItZmYxOC00ZmNlLWJjYjgtN2Y4ODJiZWJmNzZjIiwidXNlcm5hbWUiOiJhIiwiZW1haWwiOiJhQGdtYWlsLmNvbSIsInJvbGUiOiJ1c2VyIiwiaXNzIjoidmVud2V4Iiwic3ViIjoiOTU3YmM1OWItZmYxOC00ZmNlLWJjYjgtN2Y4ODJiZWJmNzZjIiwiZXhwIjoxNzc3NjUxOTU5LCJpYXQiOjE3Nzc2NTEwNTl9.cyjSiizrlBAJ-JgDqAGIF7supoQ9BriyIlT1Rrgb0cw",
//"refresh_token": "eM0q-PWAMyV67OAPHcUasah2YjfnWaNJXBQzDil75MY"
//}
