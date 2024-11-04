package main

import (
	"net/http"
	"strconv"
	"time"
)

// Assuming db is your database connection instance

type Activity struct {
	activityID        int
	title             string
	proposer          int
	startDate         time.Time
	endDate           time.Time
	maxNumber         int
	format            string
	description       string
	proposeDateTime   time.Time
	acceptAdmin       int
	acceptDateTime    time.Time
	applicationStatus string
}

type ActivityRoles struct {
	activityID   int
	activityRole string
}

func activityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getActivity(w, r)
	case "POST":
		postActivity(w, r)
	case "PATCH":
		patchActivity(w, r)
	case "DELETE":
		deleteActivity(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getActivity(w http.ResponseWriter, r *http.Request) {
	var activity Activity
	err := db.QueryRow("SELECT id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus FROM activities WHERE id = ?", activity.activityID).Scan(
		&activity.activityID, &activity.title, &activity.proposer, &activity.startDate, &activity.endDate, &activity.maxNumber, &activity.format, &activity.description, &activity.proposeDateTime, &activity.acceptAdmin, &activity.acceptDateTime, &activity.applicationStatus)
	if err != nil {
		http.Error(w, "Failed to fetch activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Add response writing logic as needed here
}

func postActivity(w http.ResponseWriter, r *http.Request) {
	activityID, _ := strconv.Atoi(r.FormValue("activityID"))
	proposer, _ := strconv.Atoi(r.FormValue("proposer"))
	maxNumber, _ := strconv.Atoi(r.FormValue("maxNumber"))
	acceptAdmin, _ := strconv.Atoi(r.FormValue("acceptAdmin"))
	startDate, _ := time.Parse("2006-01-02", r.FormValue("startDate"))
	endDate, _ := time.Parse("2006-01-02", r.FormValue("endDate"))
	proposeDateTime, _ := time.Parse(time.RFC3339, r.FormValue("proposeDateTime"))
	acceptDateTime, _ := time.Parse(time.RFC3339, r.FormValue("acceptDateTime"))

	activity := Activity{
		activityID:        activityID,
		title:             r.FormValue("title"),
		proposer:          proposer,
		startDate:         startDate,
		endDate:           endDate,
		maxNumber:         maxNumber,
		format:            r.FormValue("format"),
		description:       r.FormValue("description"),
		proposeDateTime:   proposeDateTime,
		acceptAdmin:       acceptAdmin,
		acceptDateTime:    acceptDateTime,
		applicationStatus: r.FormValue("applicationStatus"),
	}

	_, err := db.Exec("INSERT INTO activities (id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, acceptAdmin, acceptDateTime, applicationStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", activity.activityID, activity.title, activity.proposer,
		activity.startDate, activity.endDate, activity.maxNumber, activity.format, activity.description, activity.proposeDateTime, activity.acceptAdmin, activity.acceptDateTime, activity.applicationStatus)
	if err != nil {
		http.Error(w, "Failed to insert activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activity", http.StatusFound)
}

func patchActivity(w http.ResponseWriter, r *http.Request) {
	activityID, _ := strconv.Atoi(r.FormValue("activityID"))
	proposer, _ := strconv.Atoi(r.FormValue("proposer"))
	maxNumber, _ := strconv.Atoi(r.FormValue("maxNumber"))
	acceptAdmin, _ := strconv.Atoi(r.FormValue("acceptAdmin"))
	startDate, _ := time.Parse("2006-01-02", r.FormValue("startDate"))
	endDate, _ := time.Parse("2006-01-02", r.FormValue("endDate"))
	proposeDateTime, _ := time.Parse(time.RFC3339, r.FormValue("proposeDateTime"))
	acceptDateTime, _ := time.Parse(time.RFC3339, r.FormValue("acceptDateTime"))

	activity := Activity{
		activityID:        activityID,
		title:             r.FormValue("title"),
		proposer:          proposer,
		startDate:         startDate,
		endDate:           endDate,
		maxNumber:         maxNumber,
		format:            r.FormValue("format"),
		description:       r.FormValue("description"),
		proposeDateTime:   proposeDateTime,
		acceptAdmin:       acceptAdmin,
		acceptDateTime:    acceptDateTime,
		applicationStatus: r.FormValue("applicationStatus"),
	}

	_, err := db.Exec("UPDATE activities SET title = ?, proposer = ?, startDate = ?, endDate = ?, maxNumber = ?, format = ?, description = ?, proposeDateTime = ?, acceptAdmin = ?, acceptDateTime = ?, applicationStatus = ? WHERE id = ?", activity.title, activity.proposer, activity.startDate, activity.endDate, activity.maxNumber,
		activity.format, activity.description, activity.proposeDateTime, activity.acceptAdmin, activity.acceptDateTime, activity.applicationStatus, activity.activityID)
	if err != nil {
		http.Error(w, "Failed to update activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activity", http.StatusFound)
}

func deleteActivity(w http.ResponseWriter, r *http.Request) {
	activityID, _ := strconv.Atoi(r.FormValue("activityID"))

	_, err := db.Exec("DELETE FROM activities WHERE id = ?", activityID)
	if err != nil {
		http.Error(w, "Failed to delete activity: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activity", http.StatusFound)
}
