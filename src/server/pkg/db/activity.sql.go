// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: activity.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const getActivityIDByTitle = `-- name: GetActivityIDByTitle :one
SELECT activityID
FROM Activity
WHERE title = ?
`

func (q *Queries) GetActivityIDByTitle(ctx context.Context, title string) (int32, error) {
	row := q.db.QueryRowContext(ctx, getActivityIDByTitle, title)
	var activityid int32
	err := row.Scan(&activityid)
	return activityid, err
}

const insertActivity = `-- name: InsertActivity :exec
INSERT INTO Activity (title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime
) VALUES (?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertActivityParams struct {
	Title           string    `json:"title"`
	Proposer        int32     `json:"proposer"`
	Startdate       time.Time `json:"startdate"`
	Enddate         time.Time `json:"enddate"`
	Maxnumber       int32     `json:"maxnumber"`
	Format          string    `json:"format"`
	Description     string    `json:"description"`
	Proposedatetime time.Time `json:"proposedatetime"`
}

func (q *Queries) InsertActivity(ctx context.Context, arg InsertActivityParams) error {
	_, err := q.db.ExecContext(ctx, insertActivity,
		arg.Title,
		arg.Proposer,
		arg.Startdate,
		arg.Enddate,
		arg.Maxnumber,
		arg.Format,
		arg.Description,
		arg.Proposedatetime,
	)
	return err
}

const insertActivityRole = `-- name: InsertActivityRole :exec
INSERT INTO ActivityRoles (activityID, activityRole) VALUES (?, ?)
`

type InsertActivityRoleParams struct {
	Activityid   int32  `json:"activityid"`
	Activityrole string `json:"activityrole"`
}

func (q *Queries) InsertActivityRole(ctx context.Context, arg InsertActivityRoleParams) error {
	_, err := q.db.ExecContext(ctx, insertActivityRole, arg.Activityid, arg.Activityrole)
	return err
}

const insertProject = `-- name: InsertProject :exec
INSERT INTO Project (projectID, advisor) VALUES (?, ?)
`

type InsertProjectParams struct {
	Projectid int32          `json:"projectid"`
	Advisor   sql.NullString `json:"advisor"`
}

func (q *Queries) InsertProject(ctx context.Context, arg InsertProjectParams) error {
	_, err := q.db.ExecContext(ctx, insertProject, arg.Projectid, arg.Advisor)
	return err
}

const insertWorkshop = `-- name: InsertWorkshop :exec
INSERT INTO Workshop (workshopID, starttime, endtime) VALUES (?, ?, ?)
`

type InsertWorkshopParams struct {
	Workshopid int32     `json:"workshopid"`
	Starttime  time.Time `json:"starttime"`
	Endtime    time.Time `json:"endtime"`
}

func (q *Queries) InsertWorkshop(ctx context.Context, arg InsertWorkshopParams) error {
	_, err := q.db.ExecContext(ctx, insertWorkshop, arg.Workshopid, arg.Starttime, arg.Endtime)
	return err
}

const listActivity = `-- name: ListActivity :one
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL AND activityID = ?
`

func (q *Queries) ListActivity(ctx context.Context, activityid int32) (Activity, error) {
	row := q.db.QueryRowContext(ctx, listActivity, activityid)
	var i Activity
	err := row.Scan(
		&i.Activityid,
		&i.Title,
		&i.Proposer,
		&i.Startdate,
		&i.Enddate,
		&i.Maxnumber,
		&i.Format,
		&i.Description,
		&i.Proposedatetime,
		&i.Acceptadmin,
		&i.Acceptdatetime,
		&i.Applicationstatus,
	)
	return i, err
}

const listActivityRoles = `-- name: ListActivityRoles :many
SELECT activityRole FROM ActivityRoles WHERE activityID = ?
`

func (q *Queries) ListActivityRoles(ctx context.Context, activityid int32) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listActivityRoles, activityid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var activityrole string
		if err := rows.Scan(&activityrole); err != nil {
			return nil, err
		}
		items = append(items, activityrole)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRequestingActivities = `-- name: ListRequestingActivities :many
SELECT activityID, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus
FROM Activity
WHERE acceptAdmin IS NULL AND acceptDateTime IS NULL AND applicationStatus IS NULL
`

func (q *Queries) ListRequestingActivities(ctx context.Context) ([]Activity, error) {
	rows, err := q.db.QueryContext(ctx, listRequestingActivities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Activity
	for rows.Next() {
		var i Activity
		if err := rows.Scan(
			&i.Activityid,
			&i.Title,
			&i.Proposer,
			&i.Startdate,
			&i.Enddate,
			&i.Maxnumber,
			&i.Format,
			&i.Description,
			&i.Proposedatetime,
			&i.Acceptadmin,
			&i.Acceptdatetime,
			&i.Applicationstatus,
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
