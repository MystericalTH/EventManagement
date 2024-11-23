// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: chat.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const insertChat = `-- name: InsertChat :exec
INSERT INTO chatDevAd (adminID, developerID, message, datetime) 
VALUES (?, ?, ?, NOW())
`

type InsertChatParams struct {
	Adminid     sql.NullInt32 `json:"adminid"`
	Developerid sql.NullInt32 `json:"developerid"`
	Message     string        `json:"message"`
}

func (q *Queries) InsertChat(ctx context.Context, arg InsertChatParams) error {
	_, err := q.db.ExecContext(ctx, insertChat, arg.Adminid, arg.Developerid, arg.Message)
	return err
}

const listChat = `-- name: ListChat :many
SELECT 
    c.messageid, 
    a.email AS admin_email, 
    d.email AS developer_email, 
    c.message, 
    c.datetime 
FROM 
    chatDevAd c
LEFT JOIN 
    Admin a ON c.adminid = a.adminID
LEFT JOIN 
    Developer d ON c.developerid = d.developerID
`

type ListChatRow struct {
	Messageid      int32          `json:"messageid"`
	AdminEmail     sql.NullString `json:"admin_email"`
	DeveloperEmail sql.NullString `json:"developer_email"`
	Message        string         `json:"message"`
	Datetime       time.Time      `json:"datetime"`
}

func (q *Queries) ListChat(ctx context.Context) ([]ListChatRow, error) {
	rows, err := q.db.QueryContext(ctx, listChat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListChatRow
	for rows.Next() {
		var i ListChatRow
		if err := rows.Scan(
			&i.Messageid,
			&i.AdminEmail,
			&i.DeveloperEmail,
			&i.Message,
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
