-- name: InsertFeedback :exec
INSERT INTO Feedback (activityID, memberID, feedbackMessage, feedbackDateTime)
VALUES (?, ?, ?, CONVERT_TZ(NOW(), 'UTC', '+07:00'));

-- name: ListFeedbacks :many
SELECT feedbackID, activityID, Member.fname, Member.lName, feedbackMessage, feedbackDateTime
FROM Feedback
JOIN Member ON Feedback.memberID = Member.memberID
WHERE activityID = ?;

-- name: HasSubmittedFeedback :one
SELECT COUNT(*) > 0 FROM Feedback WHERE activityID = ? AND memberID = ?;