-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY ,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    token TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE users;