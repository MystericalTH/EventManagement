package services

import (
	"context"
	"sinno-server/pkg/db"
)

// Get all activities service
func GetAllActivitiesService(queries *db.Queries) ([]db.ListRequestingActivitiesRow, error) {
	return queries.ListRequestingActivities(context.Background())
}

// Get activity by ID service
func GetActivityByIDService(queries *db.Queries, activityID int32) (db.ListActivityRow, error) {
	return queries.ListActivity(context.Background(), activityID)
}

// Create activity service
func CreateActivityService(queries *db.Queries, params db.InsertActivityParams) error {
	return queries.InsertActivity(context.Background(), params)
}

// Get activity roles service
func GetActivityRolesService(queries *db.Queries, activityID int32) ([]string, error) {
	return queries.ListActivityRoles(context.Background(), activityID)
}

// Insert activity role service
func InsertActivityRoleService(queries *db.Queries, activityID int32, role string) error {
	params := db.InsertActivityRoleParams{
		Activityid:   activityID,
		Activityrole: role,
	}
	return queries.InsertActivityRole(context.Background(), params)
}

// GetActivityIDByTitleService retrieves the activity ID based on the title
func GetActivityIDByTitleService(queries *db.Queries, title string) (int32, error) {
	return queries.GetActivityIDByTitle(context.Background(), title)
}
