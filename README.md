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

авторизация, пагинация, middlewares, фронт нормальный и все будет бомба, docker compose (multi stage build)
щяс нужно подумать как запускать фронт и где его хранить. 
