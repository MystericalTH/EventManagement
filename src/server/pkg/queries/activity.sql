-- name: ListRequestingActivities :many
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL;

-- name: ListActivity :one
SELECT a.activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus, startTime, endTime, advisor, roles
  FROM Activity a 
  LEFT JOIN Workshop w ON a.activityID = w.workshopID
  LEFT JOIN Project p ON a.activityID = p.projectID
  LEFT JOIN 
    (
        SELECT 
            activityID, 
            GROUP_CONCAT(role) AS roles
        FROM 
            ActivityRoles
        GROUP BY 
            activityID
    ) ar ON a.activityID = ar.activityID
WHERE a.activityID = ?;

-- name: ListAcceptedActivities :many
SELECT a.activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus, startTime, endTime, advisor, roles
  FROM Activity a 
  LEFT JOIN Workshop w ON a.activityID = w.workshopID
  LEFT JOIN Project p ON a.activityID = p.projectID
  LEFT JOIN 
    (
        SELECT 
            activityID, 
            GROUP_CONCAT(role) AS roles
        FROM 
            ActivityRoles
        GROUP BY 
            activityID
    ) ar ON a.activityID = ar.activityID
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

