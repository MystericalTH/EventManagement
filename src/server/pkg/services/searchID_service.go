package services

import (
	"context"
	"sinno-server/pkg/db"
)

// CreateRegistrationService creates a new registration entry
func GetAdminIDByEmailService(queries *db.Queries, email string) (int32, error) {
	return queries.GetAdminIDByEmail(context.Background(), email)
}

func GetMemberIDByEmailService(queries *db.Queries, email string) (int32, error) {
	return queries.GetMemberIDByEmail(context.Background(), email)
}

func GetDeveloperIDByEmailService(queries *db.Queries, email string) (int32, error) {
	return queries.GetDeveloperIDByEmail(context.Background(), email)
}
