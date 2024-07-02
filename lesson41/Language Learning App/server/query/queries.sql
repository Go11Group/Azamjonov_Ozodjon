create database language_learning_app;

create table users(
    user_id uuid Primary Key default gen_random_uuid(),
    name varchar(100) NOT NULL,
    email varchar(100) unique NOT NULL,
    birthday timestamp NOT NULL,
    password varchar(100) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

create table courses(
    course_id uuid Primary Key default gen_random_uuid(),
    title varchar(100) NOT NULL,
    description text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

create table lessons(
    lesson_id uuid Primary Key default gen_random_uuid(),
    course_id uuid references courses(course_id) NOT NULL,
    title varchar(100) NOT NULL,
    content text NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);

create table enrollments(
    enrollment_id uuid Primary Key default gen_random_uuid(),
    user_id uuid references users(user_id) NOT NULL,
    course_id uuid references courses(course_id) NOT NULL,
    enrollment_date timestamp NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp,
    deleted_at timestamp
);