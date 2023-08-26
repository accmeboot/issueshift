-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users
ORDER BY email;

-- name: CreateUser :exec
INSERT INTO users (
    email, name, password_hash, avatar_url
) VALUES ($1, $2, $3, $4);

-- name: UpdateUser :one
UPDATE users SET name = $1, avatar_url = $2 WHERE email = $3 RETURNING name, email, avatar_url;

-- name: DeleteUser :exec
DELETE FROM users
WHERE email = $1;