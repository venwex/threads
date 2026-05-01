-- +goose Up
SELECT 'up SQL query';
create table if not exists posts (
    post_id uuid primary key default gen_random_uuid(),
    author_id uuid not null references users(user_id) on delete cascade,
    content text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);
-- +goose Down
SELECT 'down SQL query';
drop table if exists posts;
