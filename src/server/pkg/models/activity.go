package models

import "time"

type Activity struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	Proposer          int       `json:"proposer"`
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	MaxNumber         int       `json:"maxNumber"`
	Format            string    `json:"format"`
	Description       string    `json:"description"`
	ProposeDateTime   time.Time `json:"proposeDateTime"`
	AcceptAdmin       int       `json:"acceptAdmin"`
	AcceptDateTime    time.Time `json:"acceptDateTime"`
	ApplicationStatus string    `json:"applicationStatus"`
}
