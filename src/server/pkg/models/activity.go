package models

import (
	"database/sql"
	"time"
)

type Activity struct {
	ID                int32          `json:"id"`
	Title             string         `json:"title"`
	Proposer          int32          `json:"proposer"`
	StartDate         time.Time      `json:"startDate"`
	EndDate           time.Time      `json:"endDate"`
	MaxNumber         sql.NullInt32  `json:"maxNumber"`
	Format            sql.NullString `json:"format"`
	Description       string         `json:"description"`
	ProposeDateTime   time.Time      `json:"proposeDateTime"`
	AcceptAdmin       int32          `json:"acceptAdmin"`
	AcceptDateTime    time.Time      `json:"acceptDateTime"`
	ApplicationStatus string         `json:"applicationStatus"`
}
