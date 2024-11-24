// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: registration.sql

package db

import (
	"context"
	"time"
)

const checkProposer = `-- name: CheckProposer :one
SELECT COUNT(1) > 0 AS isProposer
    FROM Activity
    WHERE activityID = ? AND proposer = ?
`

type CheckProposerParams struct {
	Activityid int32 `json:"activityid"`
	Proposer   int32 `json:"proposer"`
}

func (q *Queries) CheckProposer(ctx context.Context, arg CheckProposerParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkProposer, arg.Activityid, arg.Proposer)
	var isproposer bool
	err := row.Scan(&isproposer)
	return isproposer, err
}

const getRegistrationStatus = `-- name: GetRegistrationStatus :one
SELECT COUNT(*) > 0 AS is_registered
FROM ActivityRegistration
WHERE activityID = ? AND memberID = ?
`

type GetRegistrationStatusParams struct {
	Activityid int32 `json:"activityid"`
	Memberid   int32 `json:"memberid"`
}

func (q *Queries) GetRegistrationStatus(ctx context.Context, arg GetRegistrationStatusParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, getRegistrationStatus, arg.Activityid, arg.Memberid)
	var is_registered bool
	err := row.Scan(&is_registered)
	return is_registered, err
}

const insertRegistration = `-- name: InsertRegistration :exec
INSERT INTO ActivityRegistration (activityID, memberID, role, expectation, datetime)
VALUES (?, ?, ?, ?, CONVERT_TZ(NOW(), 'UTC', '+07:00'))
`

type InsertRegistrationParams struct {
	Activityid  int32  `json:"activityid"`
	Memberid    int32  `json:"memberid"`
	Role        string `json:"role"`
	Expectation string `json:"expectation"`
}

func (q *Queries) InsertRegistration(ctx context.Context, arg InsertRegistrationParams) error {
	_, err := q.db.ExecContext(ctx, insertRegistration,
		arg.Activityid,
		arg.Memberid,
		arg.Role,
		arg.Expectation,
	)
	return err
}

const listActivityRegistration = `-- name: ListActivityRegistration :many
SELECT Member.fname, Member.lname, role, Member.email, Member.phone, expectation, datetime
FROM ActivityRegistration
JOIN Member ON ActivityRegistration.memberID = Member.memberID
WHERE activityID = ?
`

type ListActivityRegistrationRow struct {
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	Role        string    `json:"role"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Expectation string    `json:"expectation"`
	Datetime    time.Time `json:"datetime"`
}

func (q *Queries) ListActivityRegistration(ctx context.Context, activityid int32) ([]ListActivityRegistrationRow, error) {
	rows, err := q.db.QueryContext(ctx, listActivityRegistration, activityid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListActivityRegistrationRow{}
	for rows.Next() {
		var i ListActivityRegistrationRow
		if err := rows.Scan(
			&i.Fname,
			&i.Lname,
			&i.Role,
			&i.Email,
			&i.Phone,
			&i.Expectation,
			&i.Datetime,
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

const listMemberActivities = `-- name: ListMemberActivities :many
SELECT 
    a.activityID, 
    a.title, 
    a.description, 
    ar.datetime, 
    a.proposer, 
    ar.role, 
    ar.expectation
    FROM 
        ActivityRegistration ar
    JOIN 
        Activity a ON ar.activityID = a.activityID
    WHERE 
        ar.memberID = ?
`

type ListMemberActivitiesRow struct {
	Activityid  int32     `json:"activityid"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Datetime    time.Time `json:"datetime"`
	Proposer    int32     `json:"proposer"`
	Role        string    `json:"role"`
	Expectation string    `json:"expectation"`
}

func (q *Queries) ListMemberActivities(ctx context.Context, memberid int32) ([]ListMemberActivitiesRow, error) {
	rows, err := q.db.QueryContext(ctx, listMemberActivities, memberid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListMemberActivitiesRow{}
	for rows.Next() {
		var i ListMemberActivitiesRow
		if err := rows.Scan(
			&i.Activityid,
			&i.Title,
			&i.Description,
			&i.Datetime,
			&i.Proposer,
			&i.Role,
			&i.Expectation,
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

const listSubmittedMembers = `-- name: ListSubmittedMembers :many
SELECT 
    m.memberID, 
    m.fName, 
    m.lName, 
    m.email, 
    m.phone, 
    ar.role, 
    ar.expectation, 
    ar.datetime
    FROM 
        ActivityRegistration ar
    JOIN 
        Member m ON ar.memberID = m.memberID
    WHERE 
        ar.activityID = ?
`

type ListSubmittedMembersRow struct {
	Memberid    int32     `json:"memberid"`
	Fname       string    `json:"fname"`
	Lname       string    `json:"lname"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Role        string    `json:"role"`
	Expectation string    `json:"expectation"`
	Datetime    time.Time `json:"datetime"`
}

func (q *Queries) ListSubmittedMembers(ctx context.Context, activityid int32) ([]ListSubmittedMembersRow, error) {
	rows, err := q.db.QueryContext(ctx, listSubmittedMembers, activityid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListSubmittedMembersRow{}
	for rows.Next() {
		var i ListSubmittedMembersRow
		if err := rows.Scan(
			&i.Memberid,
			&i.Fname,
			&i.Lname,
			&i.Email,
			&i.Phone,
			&i.Role,
			&i.Expectation,
			&i.Datetime,
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
