-- name: ListRequestingActivities :many
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL;

-- name: ListActivity :one
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL AND activityID = ?;

-- name: InsertActivity :exec
INSERT INTO Activity (title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: ListActivityRoles :many
SELECT activityRole FROM ActivityRoles WHERE activityID = ?;

-- name: InsertActivityRole :exec
INSERT INTO ActivityRoles (activityID, activityRole) VALUES (?, ?);