package services

import (
	"context"
	"fmt"
	"log"
	"os"
	"sinno-server/pkg/db"
	secure "sinno-server/pkg/utils/dbsecure"
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
	// Fetch submitted members from the database
	members, err := queries.ListSubmittedMembers(context.Background(), activityID)
	if err != nil {
		log.Printf("Error fetching submitted members for activity ID %d: %v", activityID, err)
		return nil, err
	}

	// Iterate through the result set to decrypt sensitive fields
	for i, member := range members {
		// Decrypt first name
		decryptedFName, err := secure.DecryptFromString(member.Fname, os.Getenv("ECRYPT_KEY"))
		if err != nil {
			log.Printf("Error decrypting first name for member ID %d: %v", member.Memberid, err)
			return nil, fmt.Errorf("failed to decrypt first name for member ID %d: %v", member.Memberid, err)
		}

		// Decrypt last name
		decryptedLName, err := secure.DecryptFromString(member.Lname, os.Getenv("ECRYPT_KEY"))
		if err != nil {
			log.Printf("Error decrypting last name for member ID %d: %v", member.Memberid, err)
			return nil, fmt.Errorf("failed to decrypt last name for member ID %d: %v", member.Memberid, err)
		}

		// Decrypt phone
		decryptedPhone, err := secure.DecryptFromString(member.Phone, os.Getenv("ECRYPT_KEY"))
		if err != nil {
			log.Printf("Error decrypting phone number for member ID %d: %v", member.Memberid, err)
			return nil, fmt.Errorf("failed to decrypt phone number for member ID %d: %v", member.Memberid, err)
		}

		// Update the member struct with decrypted values
		members[i].Fname = decryptedFName
		members[i].Lname = decryptedLName
		members[i].Phone = decryptedPhone
	}

	// Log successful operation
	log.Printf("Successfully decrypted %d submitted members for activity ID %d", len(members), activityID)
	return members, nil
}

func GetEngagements(queries *db.Queries, memberID int32) ([]db.ListEngagementsRow, error) {
	return queries.ListEngagements(context.Background(), memberID)
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
