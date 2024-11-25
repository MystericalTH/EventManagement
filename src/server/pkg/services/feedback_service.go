package services

import (
	"context"
	"fmt"
	"log"
	"sinno-server/pkg/db"
	secure "sinno-server/pkg/utils/dbsecure"
)

// CreateFeedbackService creates a new feedback entry
func CreateFeedbackService(queries *db.Queries, params db.InsertFeedbackParams) error {
	return queries.InsertFeedback(context.Background(), params)
}

// GetFeedbacksByActivityService lists all feedback entries for an activity with decrypted names
func GetFeedbacksByActivityService(queries *db.Queries, activityID int32) ([]db.ListFeedbacksRow, error) {
	// Fetch feedbacks from the database
	feedbacks, err := queries.ListFeedbacks(context.Background(), activityID)
	if err != nil {
		log.Printf("Error fetching feedbacks for activity ID %d: %v", activityID, err)
		return nil, err
	}

	// Iterate through the result set to decrypt any sensitive fields (like member names)
	for i, feedback := range feedbacks {
		// Decrypt first name of the member who provided the feedback
		decryptedFName, err := secure.DecryptFromString(feedback.Fname, "database_present")
		if err != nil {

			return nil, fmt.Errorf("failed to decrypt first name for member ID")
		}

		// Decrypt last name of the member who provided the feedback
		decryptedLName, err := secure.DecryptFromString(feedback.Lname, "database_present")
		if err != nil {
			log.Printf("Error decrypting last name for member ID")
			return nil, fmt.Errorf("failed to decrypt last name for member ID")
		}

		// Update the feedback struct with decrypted values
		feedbacks[i].Fname = decryptedFName
		feedbacks[i].Lname = decryptedLName
	}

	// Log successful operation
	log.Printf("Successfully decrypted %d feedbacks for activity ID %d", len(feedbacks), activityID)
	return feedbacks, nil
}

// HasSubmittedFeedbackService checks if feedback has been submitted by a member for an activity
func HasSubmittedFeedbackService(queries *db.Queries, activityID int32, memberID int32) (bool, error) {
	params := db.HasSubmittedFeedbackParams{
		Activityid: activityID,
		Memberid:   memberID,
	}
	return queries.HasSubmittedFeedback(context.Background(), params)
}
