entities: users, posts and comments

type User struct {
ID        int       `json:"id" db:"id"`                    serial primary key
Username  string    `json:"username" db:"username"`        varchar(50) not null unique
Email     string    `json:"email" db:"email"`              varchar(50) not null unique
PasswordHash  string    `json:"-" db:"password_hash"`      varchar(255) not null
CreatedAt time.Time `json:"createdAt" db:"created_at"`     timestampz not null default now()
UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`     timestampz not null default now()
}

type Post struct {
ID        int       `json:"id" db:"id"`                     serial primary key
AuthorID  int       `json:"author_id" db:"author_id"`       references users(id)
Content   string    `json:"content" db:"content"`           Text not null
CreatedAt time.Time `json:"created_at" db:"created_at"`     timestampz not null default now()
UpdatedAt time.Time `json:"updated_at" db:"updated_at"`     timestampz not null default now()
}

авторизация, пагинация, фронт нормальный и все будет бомба, docker compose (multi stage build)
щяс нужно подумать как запускать фронт и где его хранить. 
configuration и брать значения из переменных окружения 

1. Authentication & Authorization:
   a. User registration and login using JWT-based authentication
   b. Secure token generation and verification
   c. Role‐based access control (e.g., admin, regular user)
   d. Proper password hashing and validation

2. CRUD operations:
   a. At least  three major entities/tables (with meaningful relationships:
   one‐to‐many, many‐to‐many, etc.)

3. Database & Migrations
   a. Use golang-migrate
   b. Schema with foreign keys and indexes
   c. Seed data where needed

4. Concurrency & Context
   a. At least one background worker (goroutines, channels)
   b. Context propagation, cancellation, graceful shutdown

5. API Documentation
   a. Basic documentation of endpoints (README with endpoints or
   Swagger/OpenAPI spec)

6. Testing
   a. Unit tests with testing package
   b. Coverage for critical endpoints

7. Code Organization & Best Practices
Containerization with Docker


создать таблицу refresh_tokens в дб
/auth/refresh - закончить ендпоинт, добавить логику обновления access_token через refresh_token
role-based access control, то есть только авторы собственных постов могут удалять посты. (users, admin)
auth_middleware
исправить в auth

исправить детали в целом проекте

docker compose