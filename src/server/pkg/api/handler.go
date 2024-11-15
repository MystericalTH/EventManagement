package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sinno-server/pkg/db"
	"sinno-server/pkg/models"

	"github.com/gorilla/sessions"
)

var (
	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionName  = "session-one"
)

func HandleVerifyRole(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, sessionName)
	role, ok := session.Values["role"].(string)
	if !ok {
		role = "unknown"
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"role\":\"%s\"}", role)
}

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
