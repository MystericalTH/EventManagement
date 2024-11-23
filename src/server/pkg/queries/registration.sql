-- name: InsertRegistration :exec
INSERT INTO ActivityRegistration (activityID, memberID, role, expectation, datetime)
VALUES (?, ?, ?, ?, NOW());

-- name: GetRegistrationStatus :one
SELECT COUNT(*) > 0 AS is_registered
FROM ActivityRegistration
WHERE activityID = ? AND memberID = ?;

-- name: ListActivityRegistration :many
SELECT Member.fname, Member.lname, role, Member.email, Member.phone, expectation, datetime
FROM ActivityRegistration
JOIN Member ON ActivityRegistration.memberID = Member.memberID
WHERE activityID = ?;