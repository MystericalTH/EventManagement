package typing

import (
	"sinno-server/pkg/db"
	"time"
)

type ChatChannelInfo struct {
	ID       int        `json:"id"`
	Fname    string     `json:"fname"`
	Lname    string     `json:"lname"`
	Message  *string    `json:"message"`
	Timesent *time.Time `json:"timesent"`
}

func ConvertListInitialAdminChatToDevRow(row db.ListInitialAdminChatToDevRow) ChatChannelInfo {
	var message *string
	if row.Message.Valid {
		messageStr := row.Message.String
		message = &messageStr
	}

	var timesent *time.Time
	if row.Timesent.Valid {
		timesentTime := row.Timesent.Time
		timesent = &timesentTime
	}

	return ChatChannelInfo{
		ID:       int(row.Developerid),
		Fname:    row.DeveloperFname,
		Lname:    row.DeveloperLname,
		Message:  message,
		Timesent: timesent,
	}
}

func ConvertListInitialDevChatToAdminRow(row db.ListInitialDevChatToAdminRow) ChatChannelInfo {
	var message *string
	if row.Message.Valid {
		messageStr := row.Message.String
		message = &messageStr
	}

	var timesent *time.Time
	if row.Timesent.Valid {
		timesentTime := row.Timesent.Time
		timesent = &timesentTime
	}

	return ChatChannelInfo{
		ID:       int(row.Adminid),
		Fname:    row.AdminFname,
		Lname:    row.AdminLname,
		Message:  message,
		Timesent: timesent,
	}
}
