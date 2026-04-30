-- +goose Up
create table if not exists posts (
    id uuid primary key default gen_random_uuid(),
    author_id uuid not null references users(id),
    content text not null,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now(),
    deleted_at timestamptz
);

-- +goose Down
drop table if exists posts;
