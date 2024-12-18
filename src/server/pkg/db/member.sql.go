// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: member.sql

package db

import (
	"context"
	"database/sql"
)

const acceptMember = `-- name: AcceptMember :exec
UPDATE MEMBER
SET acceptDateTime = CONVERT_TZ(NOW(), 'UTC', '+07:00'),
    acceptAdmin = ? -- Include the admin responsible for the approval
WHERE memberID = ?
`

type AcceptMemberParams struct {
	Acceptadmin sql.NullInt32 `json:"acceptadmin"`
	Memberid    int32         `json:"memberid"`
}

func (q *Queries) AcceptMember(ctx context.Context, arg AcceptMemberParams) error {
	_, err := q.db.ExecContext(ctx, acceptMember, arg.Acceptadmin, arg.Memberid)
	return err
}

const deleteMember = `-- name: DeleteMember :exec
DELETE FROM MEMBER
WHERE memberID = ?
`

func (q *Queries) DeleteMember(ctx context.Context, memberid int32) error {
	_, err := q.db.ExecContext(ctx, deleteMember, memberid)
	return err
}

const insertMember = `-- name: InsertMember :exec
INSERT INTO MEMBER (fName, lName, email, phone, githubUrl, interest, reason) 
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type InsertMemberParams struct {
	Fname     string `json:"fname"`
	Lname     string `json:"lname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Githuburl string `json:"githuburl"`
	Interest  string `json:"interest"`
	Reason    string `json:"reason"`
}

func (q *Queries) InsertMember(ctx context.Context, arg InsertMemberParams) error {
	_, err := q.db.ExecContext(ctx, insertMember,
		arg.Fname,
		arg.Lname,
		arg.Email,
		arg.Phone,
		arg.Githuburl,
		arg.Interest,
		arg.Reason,
	)
	return err
}

const listAcceptedMembers = `-- name: ListAcceptedMembers :many
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE acceptDateTime IS NOT NULL
`

type ListAcceptedMembersRow struct {
	Memberid  int32  `json:"memberid"`
	Fname     string `json:"fname"`
	Lname     string `json:"lname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Githuburl string `json:"githuburl"`
	Interest  string `json:"interest"`
	Reason    string `json:"reason"`
}

func (q *Queries) ListAcceptedMembers(ctx context.Context) ([]ListAcceptedMembersRow, error) {
	rows, err := q.db.QueryContext(ctx, listAcceptedMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAcceptedMembersRow{}
	for rows.Next() {
		var i ListAcceptedMembersRow
		if err := rows.Scan(
			&i.Memberid,
			&i.Fname,
			&i.Lname,
			&i.Email,
			&i.Phone,
			&i.Githuburl,
			&i.Interest,
			&i.Reason,
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

const listMember = `-- name: ListMember :one
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE memberID = ?
`

type ListMemberRow struct {
	Memberid  int32  `json:"memberid"`
	Fname     string `json:"fname"`
	Lname     string `json:"lname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Githuburl string `json:"githuburl"`
	Interest  string `json:"interest"`
	Reason    string `json:"reason"`
}

func (q *Queries) ListMember(ctx context.Context, memberid int32) (ListMemberRow, error) {
	row := q.db.QueryRowContext(ctx, listMember, memberid)
	var i ListMemberRow
	err := row.Scan(
		&i.Memberid,
		&i.Fname,
		&i.Lname,
		&i.Email,
		&i.Phone,
		&i.Githuburl,
		&i.Interest,
		&i.Reason,
	)
	return i, err
}

const listMemberByEmail = `-- name: ListMemberByEmail :one
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE email = ?
`

type ListMemberByEmailRow struct {
	Memberid  int32  `json:"memberid"`
	Fname     string `json:"fname"`
	Lname     string `json:"lname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Githuburl string `json:"githuburl"`
	Interest  string `json:"interest"`
	Reason    string `json:"reason"`
}

func (q *Queries) ListMemberByEmail(ctx context.Context, email string) (ListMemberByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, listMemberByEmail, email)
	var i ListMemberByEmailRow
	err := row.Scan(
		&i.Memberid,
		&i.Fname,
		&i.Lname,
		&i.Email,
		&i.Phone,
		&i.Githuburl,
		&i.Interest,
		&i.Reason,
	)
	return i, err
}

const listRequestingMembers = `-- name: ListRequestingMembers :many
SELECT memberID, fName, lName, email, phone, githubUrl, interest, reason 
FROM MEMBER
WHERE acceptDateTime IS NULL
`

type ListRequestingMembersRow struct {
	Memberid  int32  `json:"memberid"`
	Fname     string `json:"fname"`
	Lname     string `json:"lname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Githuburl string `json:"githuburl"`
	Interest  string `json:"interest"`
	Reason    string `json:"reason"`
}

func (q *Queries) ListRequestingMembers(ctx context.Context) ([]ListRequestingMembersRow, error) {
	rows, err := q.db.QueryContext(ctx, listRequestingMembers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListRequestingMembersRow{}
	for rows.Next() {
		var i ListRequestingMembersRow
		if err := rows.Scan(
			&i.Memberid,
			&i.Fname,
			&i.Lname,
			&i.Email,
			&i.Phone,
			&i.Githuburl,
			&i.Interest,
			&i.Reason,
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
