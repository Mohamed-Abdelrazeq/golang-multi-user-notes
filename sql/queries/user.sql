-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = $1;
