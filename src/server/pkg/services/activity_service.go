package services

import (
	"context"

	"database/sql"
	"sinno-server/pkg/db"
)

// Get activity ID by title
func GetActivityIDByTitleService(queries *db.Queries, title string) (int32, error) {
	return queries.GetActivityIDByTitle(context.Background(), title)
}

// Insert activity service
func InsertActivityService(queries *db.Queries, params db.InsertActivityParams) error {
	return queries.InsertActivity(context.Background(), params)
}

// Insert activity role service
func InsertActivityRoleService(queries *db.Queries, params db.InsertActivityRoleParams) error {
	return queries.InsertActivityRole(context.Background(), params)
}

// Insert project service
func InsertProjectService(queries *db.Queries, params db.InsertProjectParams) error {
	return queries.InsertProject(context.Background(), params)
}

// Insert workshop service
func InsertWorkshopService(queries *db.Queries, params db.InsertWorkshopParams) error {
	return queries.InsertWorkshop(context.Background(), params)
}

// Get activity by ID service
func GetActivityByIDService(queries *db.Queries, activityID int32) (db.ListActivityRow, error) {
	return queries.ListActivity(context.Background(), activityID)
}

// Get activity roles service
func GetActivityRolesService(queries *db.Queries, activityID int32) ([]string, error) {
	return queries.ListActivityRoles(context.Background(), activityID)
}

// Get all activities service
func GetAllActivitiesService(queries *db.Queries) ([]db.ListRequestingActivitiesRow, error) {
	return queries.ListRequestingActivities(context.Background())
}

func ApproveActivityRegistrationService(queries *db.Queries, activityID int32, adminID int32) error {
	params := db.ApproveActivityRegistrationParams{
		Acceptadmin: sql.NullInt32{Int32: adminID, Valid: true},
		Activityid:  activityID,
	}
	return queries.ApproveActivityRegistration(context.Background(), params)
}

func DeleteActivityService(queries *db.Queries, activityID int32) error {
	return queries.DeleteActivity(context.Background(), activityID)
}

func GetAcceptedActivitiesService(queries *db.Queries) ([]db.ListAcceptedActivitiesRow, error) {
	return queries.ListAcceptedActivities(context.Background())
}
