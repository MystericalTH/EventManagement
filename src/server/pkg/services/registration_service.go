package services

import (
	"context"
	"sinno-server/pkg/db"
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

// GetActivityRegistrationService retrieves all registrations for an activity
func GetActivityRegistrationService(queries *db.Queries, activityID int32) ([]db.ListActivityRegistrationRow, error) {
	return queries.ListActivityRegistration(context.Background(), activityID)
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
