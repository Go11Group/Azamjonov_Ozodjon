
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    password VARCHAR(100) NOT NULL,
    CreatedAt timestamp NOT NULL default now(),
    UpdatedAt timestamp,
    DeletedAt timestamp
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    stock_quantity INT NOT NULL,
    CreatedAt timestamp NOT NULL default now(),
    UpdatedAt timestamp,
    DeletedAt timestamp
);

CREATE TABLE user_products (
    id SERIAL PRIMARY KEY,
    user_id int references users(id) NOT NULL,
    product_id int references products(id) NOT NULL,
    CreatedAt timestamp NOT NULL default now(),
    UpdatedAt timestamp,
    DeletedAt timestamp
);
