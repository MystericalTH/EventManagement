// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feedback.sql

package db

import (
	"context"
	"time"
)

const hasSubmittedFeedback = `-- name: HasSubmittedFeedback :one
SELECT COUNT(*) > 0 FROM Feedback WHERE activityID = ? AND memberID = ?
`

type HasSubmittedFeedbackParams struct {
	Activityid int32 `json:"activityid"`
	Memberid   int32 `json:"memberid"`
}

func (q *Queries) HasSubmittedFeedback(ctx context.Context, arg HasSubmittedFeedbackParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, hasSubmittedFeedback, arg.Activityid, arg.Memberid)
	var column_1 bool
	err := row.Scan(&column_1)
	return column_1, err
}

const insertFeedback = `-- name: InsertFeedback :exec
INSERT INTO Feedback (activityID, memberID, feedbackMessage, feedbackDateTime)
VALUES (?, ?, ?, CONVERT_TZ(NOW(), 'UTC', '+07:00'))
`

type InsertFeedbackParams struct {
	Activityid      int32  `json:"activityid"`
	Memberid        int32  `json:"memberid"`
	Feedbackmessage string `json:"feedbackmessage"`
}

func (q *Queries) InsertFeedback(ctx context.Context, arg InsertFeedbackParams) error {
	_, err := q.db.ExecContext(ctx, insertFeedback, arg.Activityid, arg.Memberid, arg.Feedbackmessage)
	return err
}

const listFeedbacks = `-- name: ListFeedbacks :many
SELECT feedbackID, activityID, Member.fname, Member.lName, feedbackMessage, feedbackDateTime
FROM Feedback
JOIN Member ON Feedback.memberID = Member.memberID
WHERE activityID = ?
`

type ListFeedbacksRow struct {
	Feedbackid       int32     `json:"feedbackid"`
	Activityid       int32     `json:"activityid"`
	Fname            string    `json:"fname"`
	Lname            string    `json:"lname"`
	Feedbackmessage  string    `json:"feedbackmessage"`
	Feedbackdatetime time.Time `json:"feedbackdatetime"`
}

func (q *Queries) ListFeedbacks(ctx context.Context, activityid int32) ([]ListFeedbacksRow, error) {
	rows, err := q.db.QueryContext(ctx, listFeedbacks, activityid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListFeedbacksRow{}
	for rows.Next() {
		var i ListFeedbacksRow
		if err := rows.Scan(
			&i.Feedbackid,
			&i.Activityid,
			&i.Fname,
			&i.Lname,
			&i.Feedbackmessage,
			&i.Feedbackdatetime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
