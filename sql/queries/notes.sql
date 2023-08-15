-- name: CreateNote :one
INSERT INTO notes (
    id,
    created_at,
    updated_at,
    content
) VALUES ($1, $2, $3, $4)
RETURNING *;
