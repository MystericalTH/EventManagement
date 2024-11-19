-- name: InsertFeedback :exec
INSERT INTO Feedback (activityID, memberID, feedbackMessage, feedbackDateTime)
VALUES (?, ?, ?, NOW());

-- name: ListFeedbackByID :one
SELECT feedbackID, activityID, memberID, feedbackMessage, feedbackDateTime
FROM Feedback
WHERE feedbackID = ?;

-- name: ListFeedbacks :many
SELECT feedbackID, activityID, memberID, feedbackMessage, feedbackDateTime
FROM Feedback;

-- name: HasSubmittedFeedback :one
SELECT COUNT(*) > 0 FROM Feedback WHERE activityID = ? AND memberID = ?;