package services

import (
	"context"
	"sinno-server/pkg/db"
)

// CreateChatService creates a new chat entry
func InsertAdminDevChatService(queries *db.Queries, params db.InsertAdminDevChatParams) error {
	return queries.InsertAdminDevChat(context.Background(), params)
}

// ListAdminDevChatService lists all chat entries between an admin and a developer
func ListAdminDevChatService(queries *db.Queries, params db.ListAdminDevChatParams) ([]db.ListAdminDevChatRow, error) {
	return queries.ListAdminDevChat(context.Background(), params)
}

// ListInitialDevChatToAdminService lists the latest chat entries from developers to an admin
func ListInitialDevChatToAdminService(queries *db.Queries, developerid int32) ([]db.ListInitialDevChatToAdminRow, error) {
	return queries.ListInitialDevChatToAdmin(context.Background(), developerid)
}

// ListInitialAdminChatToDevService lists the latest chat entries from admins to a developer
func ListInitialAdminChatToDevService(queries *db.Queries, adminid int32) ([]db.ListInitialAdminChatToDevRow, error) {
	return queries.ListInitialAdminChatToDev(context.Background(), adminid)
}
