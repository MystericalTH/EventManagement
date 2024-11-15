package api

import (
	"encoding/json"
	"net/http"
	"sinno-server/pkg/db"
	"sinno-server/pkg/models"
)

func GetActivities(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT * FROM Activity")
	if err != nil {
		http.Error(w, "Failed to fetch activities", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var activities []models.Activity
	for rows.Next() {
		var activity models.Activity
		if err := rows.Scan(
			&activity.ID,
			&activity.Title,
			&activity.Proposer,
			&activity.StartDate,
			&activity.EndDate,
			&activity.MaxNumber,
			&activity.Format,
			&activity.Description,
			&activity.ProposeDateTime,
			&activity.AcceptAdmin,
			&activity.AcceptDateTime,
			&activity.ApplicationStatus,
		); err != nil {
			http.Error(w, "Failed to scan activity", http.StatusInternalServerError)
			return
		}
		activities = append(activities, activity)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(activities)
}
