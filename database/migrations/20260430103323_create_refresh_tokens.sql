-- +goose Up
create table if not exists refresh_tokens (
    id uuid primary key default gen_random_uuid(),
    user_id uuid references users(id) on delete cascade,
    refresh_hash text not null unique,
    created_at timestamptz not null default now(),
    expires_at timestamptz not null,
    revoked_at timestamptz
);
-- +goose Down
drop table if exists refresh_tokens;