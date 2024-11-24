package typing

import (
	"sinno-server/pkg/db"
	"time"
)

type DevChannelInitialMessage struct {
	DeveloperID    int        `json:"developerID"`
	DeveloperFname string     `json:"developer_fname"`
	DeveloperLname string     `json:"developer_lname"`
	Message        *string    `json:"message"`
	Timesent       *time.Time `json:"timesent"`
}

func ConvertListInitialAdminChatToDevRow(row db.ListInitialAdminChatToDevRow) DevChannelInitialMessage {
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

	return DevChannelInitialMessage{
		DeveloperID:    int(row.Developerid),
		DeveloperFname: row.DeveloperFname,
		DeveloperLname: row.DeveloperLname,
		Message:        message,
		Timesent:       timesent,
	}
}

type AdminChannelInitialMessage struct {
	AdminID    int        `json:"adminID"`
	AdminFname string     `json:"admin_fname"`
	AdminLname string     `json:"admin_lname"`
	Message    *string    `json:"message"`
	Timesent   *time.Time `json:"timesent"`
}

func ConvertListInitialDevChatToAdminRow(row db.ListInitialDevChatToAdminRow) AdminChannelInitialMessage {
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

	return AdminChannelInitialMessage{
		AdminID:    int(row.Adminid),
		AdminFname: row.AdminFname,
		AdminLname: row.AdminLname,
		Message:    message,
		Timesent:   timesent,
	}
}
