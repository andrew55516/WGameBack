-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: CheckUserByEmail :one
SELECT EXISTS(SELECT * FROM users WHERE email = $1);

-- name: CreateUser :exec
INSERT INTO users (user_id, user_name, password, email, user_type, token, refresh_token, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);
