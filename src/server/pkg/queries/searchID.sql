-- name: GetAdminIDByEmail :one
SELECT adminID FROM Admin WHERE email = ?;

-- name: GetMemberIDByEmail :one
SELECT memberID FROM Member WHERE email = ?;

-- name: GetDeveloperIDByEmail :one
SELECT developerID FROM Developer WHERE email = ?;