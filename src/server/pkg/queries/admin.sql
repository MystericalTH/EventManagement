-- name: FetchAdminIDByEmail :one
SELECT adminID
FROM admin
WHERE email = ?;