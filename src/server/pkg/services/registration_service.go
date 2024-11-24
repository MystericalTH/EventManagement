package services

import (
	"context"
	"sinno-server/pkg/db"
	"time"
)

// CreateRegistrationService creates a new registration entry
func CreateRegistrationService(queries *db.Queries, params db.InsertRegistrationParams) error {
	return queries.InsertRegistration(context.Background(), params)
}

// GetRegistrationStatusService checks if a member is registered for an activity
func GetRegistrationStatusService(queries *db.Queries, activityID int32, memberID int32) (bool, error) {
	params := db.GetRegistrationStatusParams{
		Activityid: activityID,
		Memberid:   memberID,
	}
	return queries.GetRegistrationStatus(context.Background(), params)
}

func CheckProposerService(queries *db.Queries, activityID, memberID int32) (bool, error) {
	result, err := queries.CheckProposer(context.Background(), db.CheckProposerParams{
		Activityid: activityID,
		Proposer:   memberID,
	})
	if err != nil {
		return false, err
	}
	return result, nil
}

func GetSubmittedMembersService(queries *db.Queries, activityID int32) ([]db.ListSubmittedMembersRow, error) {
	return queries.ListSubmittedMembers(context.Background(), activityID)
}

func GetMemberActivitiesService(queries *db.Queries, memberID int32) ([]db.ListMemberActivitiesRow, error) {
	return queries.ListMemberActivities(context.Background(), memberID)
}

func CheckProjectDateConflict(queries *db.Queries, startDate, endDate time.Time) (bool, error) {
	conflictCount, err := queries.CheckProjectDateConflict(context.Background(), db.CheckProjectDateConflictParams{
		Startdate: startDate,
		Enddate:   endDate,
	})
	if err != nil {
		return false, err // Return false with the error
	}
	// If conflictCount > 0, there is a conflict
	return conflictCount > 0, nil
}

func CheckWorkshopConflict(queries *db.Queries, startDate, endDate, startTime, endTime time.Time) (bool, error) {
	conflictCount, err := queries.CheckWorkshopDateConflict(context.Background(), db.CheckWorkshopDateConflictParams{
		Startdate: startDate,
		Enddate:   endDate,
		Starttime: startTime.Format("15:04"), // Convert to HH:mm format
		Endtime:   endTime.Format("15:04"),   // Convert to HH:mm format
	})
	if err != nil {
		return false, err // Return false with the error
	}
	// If conflictCount > 0, there is a conflict
	return conflictCount > 0, nil
}
