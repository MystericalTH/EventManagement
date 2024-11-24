package typing

import (
	"fmt"
	"sinno-server/pkg/db"
	"strings"
	"time"
)

type Activity struct {
	ID                int      `json:"id"`
	Title             string   `json:"title"`
	Proposer          string   `json:"proposer"`
	StartDate         string   `json:"startDate"`
	StartTime         *string  `json:"startTime,omitempty"`
	EndDate           string   `json:"endDate"`
	EndTime           *string  `json:"endTime,omitempty"`
	MaxParticipant    int      `json:"maxParticipant"`
	Format            string   `json:"format"`
	Description       string   `json:"description"`
	Roles             []string `json:"roles"`
	ProposeDateTime   string   `json:"proposeDateTime"`
	AcceptAdmin       *string  `json:"acceptAdmin,omitempty"`
	AcceptDateTime    *string  `json:"acceptDateTime,omitempty"`
	ApplicationStatus *string  `json:"applicationStatus,omitempty"`
	Advisor           *string  `json:"advisor,omitempty"`
}

func ConvertToActivity(row db.ListActivityRow) (Activity, error) {
	var startTime, endTime, acceptAdmin, acceptDateTime, applicationStatus, advisor *string

	if row.Starttime.Valid {
		startTime = &row.Starttime.String
	}
	if row.Endtime.Valid {
		endTime = &row.Endtime.String
	}
	if row.Acceptadmin.Valid {
		acceptAdminStr := fmt.Sprintf("%d", row.Acceptadmin.Int32)
		acceptAdmin = &acceptAdminStr
	}
	if row.Acceptdatetime.Valid {
		acceptDateTimeStr := row.Acceptdatetime.Time.Format(time.RFC3339)
		acceptDateTime = &acceptDateTimeStr
	}
	if row.Applicationstatus.Valid {
		applicationStatus = &row.Applicationstatus.String
	}
	if row.Advisor.Valid {
		advisor = &row.Advisor.String
	}
	var rolesArray []string
	if row.Roles.Valid {
		rolesArray = strings.Split(row.Roles.String, ",")
	}

	return Activity{
		ID:                int(row.Activityid),
		Title:             row.Title,
		Proposer:          fmt.Sprintf("%d", row.Proposer),
		StartDate:         row.Startdate.Format("2006-01-02"),
		StartTime:         startTime,
		EndDate:           row.Enddate.Format("2006-01-02"),
		EndTime:           endTime,
		MaxParticipant:    int(row.Maxnumber),
		Format:            row.Format,
		Description:       row.Description,
		Roles:             rolesArray,
		ProposeDateTime:   row.Proposedatetime.Format(time.RFC3339),
		AcceptAdmin:       acceptAdmin,
		AcceptDateTime:    acceptDateTime,
		ApplicationStatus: applicationStatus,
		Advisor:           advisor,
	}, nil
}
