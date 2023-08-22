-- +goose Up
ALTER TABLE users
DROP COLUMN token;

-- +goose Down
ALTER TABLE users
ADD COLUMN token TEXT NOT NULL UNIQUE;