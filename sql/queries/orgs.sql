-- name: Get all organizations
INSERT INTO organizations (metamask_address, acad_email, org_name, registration_number, business_type, verification_status)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


