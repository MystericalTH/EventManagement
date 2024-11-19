package services

import (
	"context"
	"sinno-server/pkg/db"
)

// FetchAdminIDService retrieves the admin ID based on the provided email.
func FetchAdminIDService(queries *db.Queries, email string) (int32, error) {
	adminID, err := queries.FetchAdminIDByEmail(context.Background(), email)
	if err != nil {
		return 0, err
	}
	return adminID, nil
}
