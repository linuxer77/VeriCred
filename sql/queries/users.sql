-- name: Get all users
INSERT INTO users (metamask_address, email, created_at, first_name, last_name, password)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;