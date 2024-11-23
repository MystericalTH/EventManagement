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

const approveActivityRegistration = `-- name: ApproveActivityRegistration :exec
UPDATE Activity
SET acceptDateTime = LOCALTIME(),
    acceptAdmin = ? -- Include the admin responsible for the approval
WHERE activityID = ?
`

type ApproveActivityRegistrationParams struct {
	Acceptadmin sql.NullInt32 `json:"acceptadmin"`
	Activityid  int32         `json:"activityid"`
}

func (q *Queries) ApproveActivityRegistration(ctx context.Context, arg ApproveActivityRegistrationParams) error {
	_, err := q.db.ExecContext(ctx, approveActivityRegistration, arg.Acceptadmin, arg.Activityid)
	return err
}

const deleteActivity = `-- name: DeleteActivity :exec
DELETE FROM Activity
WHERE ActivityID = ?
`

func (q *Queries) DeleteActivity(ctx context.Context, activityid int32) error {
	_, err := q.db.ExecContext(ctx, deleteActivity, activityid)
	return err
}

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
	Workshopid int32  `json:"workshopid"`
	Starttime  string `json:"starttime"`
	Endtime    string `json:"endtime"`
}

func (q *Queries) InsertWorkshop(ctx context.Context, arg InsertWorkshopParams) error {
	_, err := q.db.ExecContext(ctx, insertWorkshop, arg.Workshopid, arg.Starttime, arg.Endtime)
	return err
}

const listAcceptedActivities = `-- name: ListAcceptedActivities :many
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
WHERE acceptAdmin IS NOT NULL AND acceptDateTime IS NOT NULL AND applicationStatus IS NOT NULL
`

type ListAcceptedActivitiesRow struct {
	Activityid        int32          `json:"activityid"`
	Title             string         `json:"title"`
	Proposer          int32          `json:"proposer"`
	Startdate         time.Time      `json:"startdate"`
	Enddate           time.Time      `json:"enddate"`
	Maxnumber         int32          `json:"maxnumber"`
	Format            string         `json:"format"`
	Description       string         `json:"description"`
	Proposedatetime   time.Time      `json:"proposedatetime"`
	Acceptadmin       sql.NullInt32  `json:"acceptadmin"`
	Acceptdatetime    sql.NullTime   `json:"acceptdatetime"`
	Applicationstatus sql.NullString `json:"applicationstatus"`
	Starttime         sql.NullString `json:"starttime"`
	Endtime           sql.NullString `json:"endtime"`
	Advisor           sql.NullString `json:"advisor"`
	Roles             sql.NullString `json:"roles"`
}

func (q *Queries) ListAcceptedActivities(ctx context.Context) ([]ListAcceptedActivitiesRow, error) {
	rows, err := q.db.QueryContext(ctx, listAcceptedActivities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAcceptedActivitiesRow
	for rows.Next() {
		var i ListAcceptedActivitiesRow
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
			&i.Starttime,
			&i.Endtime,
			&i.Advisor,
			&i.Roles,
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

const listActivity = `-- name: ListActivity :one
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
WHERE a.activityID = ?
`

type ListActivityRow struct {
	Activityid        int32          `json:"activityid"`
	Title             string         `json:"title"`
	Proposer          int32          `json:"proposer"`
	Startdate         time.Time      `json:"startdate"`
	Enddate           time.Time      `json:"enddate"`
	Maxnumber         int32          `json:"maxnumber"`
	Format            string         `json:"format"`
	Description       string         `json:"description"`
	Proposedatetime   time.Time      `json:"proposedatetime"`
	Acceptadmin       sql.NullInt32  `json:"acceptadmin"`
	Acceptdatetime    sql.NullTime   `json:"acceptdatetime"`
	Applicationstatus sql.NullString `json:"applicationstatus"`
	Starttime         sql.NullString `json:"starttime"`
	Endtime           sql.NullString `json:"endtime"`
	Advisor           sql.NullString `json:"advisor"`
	Roles             sql.NullString `json:"roles"`
}

func (q *Queries) ListActivity(ctx context.Context, activityid int32) (ListActivityRow, error) {
	row := q.db.QueryRowContext(ctx, listActivity, activityid)
	var i ListActivityRow
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
		&i.Starttime,
		&i.Endtime,
		&i.Advisor,
		&i.Roles,
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
