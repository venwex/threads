create table if not exists posts (
    id serial primary key,
    author_id int not null references users(id),
    content text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);