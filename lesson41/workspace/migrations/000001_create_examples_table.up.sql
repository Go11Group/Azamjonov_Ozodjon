create table examples(
    id uuid primary key default gen_random_uuid() not null,
    name varchar not null,
    created_at timestamp not null default now(),
    update_at timestamp,
    deleted_at timestamp
);