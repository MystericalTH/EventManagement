-- name: ListRequestingActivities :many
SELECT a.activityID, title, proposer, startDate, endDate, maxParticipant, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus, startTime, endTime, advisor, roles
  FROM Activity a 
  LEFT JOIN Workshop w ON a.activityID = w.workshopID
  LEFT JOIN Project p ON a.activityID = p.projectID
  LEFT JOIN 
    (
        SELECT 
            activityID, 
            GROUP_CONCAT(activityRole) AS roles
        FROM 
            ActivityRoles
        GROUP BY 
            activityID
    ) ar ON a.activityID = ar.activityID
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL;

-- name: ListActivity :one
SELECT a.activityID, title, proposer, startDate, endDate, maxParticipant, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus, startTime, endTime, advisor, roles
  FROM Activity a 
  LEFT JOIN Workshop w ON a.activityID = w.workshopID
  LEFT JOIN Project p ON a.activityID = p.projectID
  LEFT JOIN 
    (
        SELECT 
            activityID, 
            GROUP_CONCAT(activityRole) AS roles
        FROM 
            ActivityRoles
        GROUP BY 
            activityID
    ) ar ON a.activityID = ar.activityID
WHERE a.activityID = ?;

-- name: ListAcceptedActivities :many
SELECT a.activityID, title, proposer, startDate, endDate, maxParticipant, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus, startTime, endTime, advisor, roles
  FROM Activity a 
  LEFT JOIN Workshop w ON a.activityID = w.workshopID
  LEFT JOIN Project p ON a.activityID = p.projectID
  LEFT JOIN 
    (
        SELECT 
            activityID, 
            group_concat(activityRole) AS roles
        FROM 
            ActivityRoles
        GROUP BY 
            activityID
    ) ar ON a.activityID = ar.activityID
WHERE acceptAdmin IS NOT NULL AND acceptDateTime IS NOT NULL AND applicationStatus IS NOT NULL;

-- name: InsertActivity :exec
INSERT INTO Activity (title, proposer, startDate, endDate, maxParticipant, format, description, proposeDateTime, applicationStatus
) VALUES (?, ?, ?, ?, ?, ?, ?, CONVERT_TZ(NOW(), 'UTC', '+07:00'), "pending");

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
SET acceptDateTime = CONVERT_TZ(NOW(), 'UTC', '+07:00'), -- Store acceptDateTime in GMT+07:00
    acceptAdmin = ?, -- Include the admin responsible for the approval
    applicationStatus = "approved"
WHERE activityID = ?;


-- name: ListAllProposedActivity :many
SELECT a.activityID, title, proposer, startDate, endDate, maxParticipant, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus, startTime, endTime, advisor, roles
  FROM Activity a 
  LEFT JOIN Workshop w ON a.activityID = w.workshopID
  LEFT JOIN Project p ON a.activityID = p.projectID
  LEFT JOIN 
    (
        SELECT 
            activityID, 
            GROUP_CONCAT(activityRole) AS roles
        FROM 
            ActivityRoles
        GROUP BY 
            activityID
    ) ar ON a.activityID = ar.activityID
WHERE proposer = ?;

-- name: CheckProjectDateConflict :one
SELECT COUNT(1)
FROM Activity a
WHERE a.format = 'project' AND
      a.startDate = ? AND
      a.endDate = ?;


-- name: CheckWorkshopDateConflict :one
SELECT COUNT(1)
FROM Workshop w
JOIN Activity a ON w.workshopID = a.activityID
WHERE a.format = 'workshop' AND
      a.startDate = ? AND
      a.endDate = ? AND
      ((w.startTime < ? AND w.endTime > ?) OR (a.startDate < ? AND a.endDate > ?));
