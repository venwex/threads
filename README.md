<img width="1920" height="1080" alt="image" src="https://github.com/user-attachments/assets/706e9723-c83e-437b-baa4-da13faf4a16b" />

# Threads

A minimal Threads-like social feed built with Go, PostgreSQL, JWT authentication and WebSocket-based realtime updates.

The project is focused on clean backend architecture, authentication flow, database migrations, soft deletes, and realtime post broadcasting without frontend frameworks or bundlers.

---

## Overview

`threads` is a small social feed application where users can register, sign in, create posts, and receive new posts in realtime through WebSocket connections.

The backend is written in Go using the standard `net/http` package with a layered architecture:

- handlers
- services
- repositories
- middleware
- PostgreSQL storage
- JWT authentication
- WebSocket hub

The frontend is a lightweight static client served directly by the Go server.

No React. No npm. No bundler. Humanity survives another day.

---

## Features

### Authentication

- User registration
- User login
- Password hashing with bcrypt
- JWT access tokens
- Refresh token generation and rotation
- Refresh tokens stored as SHA-256 hashes
- Protected routes using authentication middleware
- User claims stored in request context

### Posts

- Create posts
- Get posts feed
- Update posts
- Delete posts
- Soft delete support with `deleted_at`
- Author information included in feed responses

### Realtime

- WebSocket endpoint for authenticated users
- Access token passed through query parameter
- Connected clients receive newly created posts instantly
- Hub-based client registration, unregistration, and broadcasting

### Database

- PostgreSQL
- UUID primary keys
- Goose migrations
- Foreign key relationships between users and posts
- Soft delete fields
- Refresh token persistence

### Frontend

- Static HTML/CSS/JavaScript client
- Login page
- Registration page
- Feed page
- Dark Threads-like UI
- Access token stored on the client
- WebSocket connection after login

---

## Tech Stack

### Backend

- Go
- net/http
- PostgreSQL
- sqlx
- goose
- bcrypt
- JWT
- gorilla/websocket
- Docker / Docker Compose

### Frontend

- HTML
- CSS
- Vanilla JavaScript

---

## Project Structure

```text
threads/
├── cmd/
│   └── app/
│       └── main.go
├── database/
│   └── migrations/
├── internal/
│   ├── auth/
│   ├── config/
│   ├── handler/
│   ├── middleware/
│   ├── models/
│   ├── repository/
│   ├── service/
│   ├── transport/
│   └── websocket/
├── web/
│   ├── index.html
│   ├── app.html
│   ├── css/
│   └── js/
├── docker-compose.yml
├── Dockerfile
├── go.mod
└── README.md
