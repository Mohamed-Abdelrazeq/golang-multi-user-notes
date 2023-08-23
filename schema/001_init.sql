-- +goose Up
CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "email" text UNIQUE NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamp
);

CREATE TABLE "notes" (
  "id" integer PRIMARY KEY,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "user_id" integer NOT NULL,
  "is_favourite" boolean DEFAULT false,
  "created_at" timestamp DEFAULT 'now()'
);

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose Down
DROP TABLE notes;
DROP TABLE users;