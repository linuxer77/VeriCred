-- name: Get all users
INSERT INTO users (metamask_address, email, first_name, last_name)
VALUES ($1, $2, $3, $4)
RETURNING *;