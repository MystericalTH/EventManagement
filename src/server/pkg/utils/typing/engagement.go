package typing

import (
	"fmt"
	"sinno-server/pkg/db"
	"time"
)

type Engagement struct {
	ActivityID        int     `json:"activityid"`
	Title             string  `json:"title"`
	StartDate         string  `json:"startdate"`
	EndDate           string  `json:"enddate"`
	StartTime         *string `json:"starttime,omitempty"`
	EndTime           *string `json:"endtime,omitempty"`
	MaxParticipant    int     `json:"maxparticipant"`
	Format            string  `json:"format"`
	Description       string  `json:"description"`
	ProposeDateTime   string  `json:"proposedatetime"`
	AcceptAdmin       *string `json:"acceptadmin,omitempty"`
	AcceptDateTime    *string `json:"acceptdatetime,omitempty"`
	ApplicationStatus *string `json:"applicationstatus,omitempty"`
	Advisor           *string `json:"advisor,omitempty"`
	Role              string  `json:"role"`
	Expectation       string  `json:"expectation"`
	EngagedAt         string  `json:"datetime"`
}

func ConvertToEngagement(row db.ListEngagementsRow) (Engagement, error) {
	var startTime, endTime, acceptAdmin, applicationStatus, acceptDateTime, advisor *string

	// Parse StartTime
	if row.Starttime.Valid {
		startTimeStr := row.Starttime.String
		parsedTime, err := time.Parse("15:04", startTimeStr)
		if err == nil {
			formattedTime := parsedTime.Format("15:04")
			startTime = &formattedTime
		} else {
			startTime = &startTimeStr
		}
	}

	// Parse EndTime
	if row.Endtime.Valid {
		endTimeStr := row.Endtime.String
		parsedTime, err := time.Parse("15:04", endTimeStr)
		if err == nil {
			formattedTime := parsedTime.Format("15:04")
			endTime = &formattedTime
		} else {
			endTime = &endTimeStr
		}
	}

	// Parse AcceptAdmin
	if row.Acceptadmin.Valid {
		acceptAdminStr := fmt.Sprintf("%d", row.Acceptadmin.Int32)
		acceptAdmin = &acceptAdminStr
	}

	// Parse AcceptDateTime
	if row.Acceptdatetime.Valid {
		acceptDateTimeStr := row.Acceptdatetime.Time.Format(time.RFC3339)
		acceptDateTime = &acceptDateTimeStr
	}

	// Parse ApplicationStatus
	if row.Applicationstatus.Valid {
		applicationStatus = &row.Applicationstatus.String
	}

	// Parse Advisor
	if row.Advisor.Valid {
		advisor = &row.Advisor.String
	}

	// Convert and return Engagement
	return Engagement{
		ActivityID:        int(row.Activityid),
		Title:             row.Title,
		StartDate:         row.Startdate.Format("2006-01-02"),
		EndDate:           row.Enddate.Format("2006-01-02"),
		StartTime:         startTime,
		EndTime:           endTime,
		MaxParticipant:    int(row.Maxparticipant),
		Format:            row.Format,
		Description:       row.Description,
		ProposeDateTime:   row.Proposedatetime.Format(time.RFC3339),
		AcceptAdmin:       acceptAdmin,
		AcceptDateTime:    acceptDateTime,
		ApplicationStatus: applicationStatus,
		Advisor:           advisor,
		Role:              row.Role,                          // Mapping "role"
		Expectation:       row.Expectation,                   // Mapping "expectation"
		EngagedAt:         row.Datetime.Format(time.RFC3339), // Mapping "datetime"
	}, nil
}
