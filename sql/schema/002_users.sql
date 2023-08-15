-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    api_key VARCHAR(64) NOT NULL UNIQUE 
);
-- +goose Down
DROP TABLE users;