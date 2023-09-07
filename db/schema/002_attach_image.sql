-- +goose Up
ALTER TABLE "notes" ADD COLUMN "image_url" TEXT NOT NULL DEFAULT '';

-- +goose Down
ALTER TABLE "notes" DROP COLUMN "image_url";