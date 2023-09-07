-- +goose Up
CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "email" text UNIQUE NOT NULL,
  "password" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);

CREATE TABLE "notes" (
  "id" SERIAL PRIMARY KEY,
  "title" text NOT NULL,
  "content" text NOT NULL,
  "user_id" integer NOT NULL,
  "is_favourite" boolean NOT NULL DEFAULT false,
  "created_at" timestamp NOT NULL DEFAULT NOW()
);

ALTER TABLE "notes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

-- +goose Down
DROP TABLE notes;
DROP TABLE users;