-- name: ListRequestingActivities :many
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL;

-- name: ListActivity :one
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL AND activityID = ?;

-- name: ListAcceptedActivities :many
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NOT NULL AND acceptDateTime IS NOT NULL AND applicationStatus IS NOT NULL;

-- name: InsertActivity :exec
INSERT INTO Activity (title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime
) VALUES (?, ?, ?, ?, ?, ?, ?, ?);

-- name: InsertProject :exec
INSERT INTO Project (projectID, advisor) VALUES (?, ?);

-- name: InsertWorkshop :exec
INSERT INTO Workshop (workshopID, starttime, endtime) VALUES (?, ?, ?);

-- name: ListActivityRoles :many
SELECT activityRole FROM ActivityRoles WHERE activityID = ?;

-- name: InsertActivityRole :exec
INSERT INTO ActivityRoles (activityID, activityRole) VALUES (?, ?);

-- name: GetActivityIDByTitle :one
SELECT activityID
FROM Activity
WHERE title = ?;

-- name: DeleteActivity :exec
DELETE FROM Activity
WHERE ActivityID = ?;

-- name: ApproveActivityRegistration :exec
UPDATE Activity
SET acceptDateTime = LOCALTIME(),
    acceptAdmin = ? -- Include the admin responsible for the approval
WHERE activityID = ?;

-- name: ListActivitiesByProposer :many
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL AND proposer = ?;