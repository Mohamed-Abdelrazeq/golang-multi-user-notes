-- +goose Up
CREATE TABLE notes (
    id UUID PRIMARY KEY NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    content TEXT NOT NULL
);

-- +goose Down
DROP TABLE notes;