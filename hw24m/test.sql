create table cars(
    id uuid primary key not null default gen_random_uuid(),
    name varchar,
    year int,
    brend varchar,
);


create table user(
    id uuid primary key not null default gen_random_uuid(),
    name varchar,
    car uuid not null references cars(id),
);