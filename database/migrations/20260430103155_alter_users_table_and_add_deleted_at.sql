-- +goose Up
SELECT 'up SQL query';
CREATE TABLE if not exists users (
    id uuid primary key default gen_random_uuid(),
    username text not null unique,
    email text not null unique,
    password text not null,
    role text not null default 'user',
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);
-- +goose Down
SELECT 'down SQL query';
drop table if exists users;