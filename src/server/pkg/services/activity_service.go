package services

import (
	"context"
	"database/sql"
	"time"

	"sinno-server/pkg/db"
)

// Get activity ID by title
func GetActivityIDByTitleService(queries *db.Queries, title string) (int32, error) {
	return queries.GetActivityIDByTitle(context.Background(), title)
}

// Insert activity service
func InsertActivityService(queries *db.Queries, title string, proposer int32, startDate time.Time, endDate time.Time, maxNumber int32, format string, description string, proposeDateTime time.Time) error {
	params := db.InsertActivityParams{
		Title:           title,
		Proposer:        proposer,
		Startdate:       startDate,
		Enddate:         endDate,
		Maxnumber:       maxNumber,
		Format:          format,
		Description:     description,
		Proposedatetime: proposeDateTime,
	}
	return queries.InsertActivity(context.Background(), params)
}

// Insert activity role service
func InsertActivityRoleService(queries *db.Queries, activityID int32, role string) error {
	params := db.InsertActivityRoleParams{
		Activityid:   activityID,
		Activityrole: role,
	}
	return queries.InsertActivityRole(context.Background(), params)
}

// Insert project service
func InsertProjectService(queries *db.Queries, advisor string) error {
	params := db.InsertProjectParams{
		Advisor: sql.NullString{String: advisor, Valid: advisor != ""},
	}
	return queries.InsertProject(context.Background(), params)
}

// Insert workshop service
func InsertWorkshopService(queries *db.Queries, startTime time.Time, endTime time.Time) error {
	params := db.InsertWorkshopParams{
		Starttime: startTime,
		Endtime:   endTime,
	}
	return queries.InsertWorkshop(context.Background(), params)
}

// Get activity by ID service
func GetActivityByIDService(queries *db.Queries, activityID int32) (db.Activity, error) {
	return queries.ListActivity(context.Background(), activityID)
}

// Get activity roles service
func GetActivityRolesService(queries *db.Queries, activityID int32) ([]string, error) {
	return queries.ListActivityRoles(context.Background(), activityID)
}

// Get all activities service
func GetAllActivitiesService(queries *db.Queries) ([]db.Activity, error) {
	return queries.ListRequestingActivities(context.Background())
}
