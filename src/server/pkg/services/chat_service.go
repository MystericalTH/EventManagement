package services

import (
	"context"
	"sinno-server/pkg/db"
)

// CreateChatService creates a new chat entry
func CreateChatService(queries *db.Queries, params db.InsertChatParams) error {
	return queries.InsertChat(context.Background(), params)
}

// ListChatsService lists all chat entries
func GetChatsService(queries *db.Queries) ([]db.ListChatRow, error) {
	return queries.ListChat(context.Background())
}
