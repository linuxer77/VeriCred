-- name: Get all accounts
INSERT INTO accounts (metamask_address, account_type, created_at, password, is_active)
VALUES ($1, $2, $3, $4, $5);
Returning *;



