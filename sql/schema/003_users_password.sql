-- +goose Up
ALTER TABLE users ADD COLUMN password_string VARCHAR(64) NOT NULL UNIQUE DEFAULT (
    "123456"
);
-- +goose Down
ALTER TABLE users DROP COLUMN password_string;