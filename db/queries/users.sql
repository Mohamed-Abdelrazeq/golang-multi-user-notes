-- name: CreateUser :one
INSERT INTO users (email, password) VALUES ($1, $2) RETURNING (id, email);

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = $1;

