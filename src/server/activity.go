package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Activity struct {
	activityID int(11),
	title varchar(255),
	proposer int(11),
	startDate date,
	endDate date,
	maxNumber int(11),
	format varchar(10),
	description text,
	proposeDateTime datetime,
	acceptAdmin int(11),
	acceptDateTime datetime,
	applicationStatus varchar(20)
}

type ActivityRoles struct {
	activityID int(11),
	activityRole varchar(30)
}

var (
	db *sql.DB
)

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
	// Query the database for activity information
	var activity Activity
	err := db.QueryRow("SELECT id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime, 
			acceptAdmin, acceptDateTime, applicationStatus FROM activities WHERE id = ?", activity.activityID).Scan(&activity.activityID, &activity.title, &activity.proposer, &activity.startDate,
			&activity.endDate, &activity.maxNumber, &activity.format, &activity.description, &activity.proposeDateTime, &activity.acceptAdmin, &activity.acceptDateTime, &activity.applicationStatus)
	if err != nil {
		http.Error(w, "Failed to fetch activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func postActivity(w http.ResponseWriter, r *http.Request) {
	// Get activity information from the request
	activity := Activity{
		activityID: r.FormValue("activityID"),
		title: r.FormValue("title"),
		proposer: r.FormValue("proposer"),
		startDate: r.FormValue("startDate"),
		endDate: r.FormValue("endDate"),
		maxNumber: r.FormValue("maxNumber"),
		format: r.FormValue("format"),
		description: r.FormValue("description"),
		proposeDateTime: r.FormValue("proposeDateTime"),
		acceptAdmin: r.FormValue("acceptAdmin"),
		acceptDateTime: r.FormValue("acceptDateTime"),
		applicationStatus: r.FormValue("applicationStatus")
	}
	
	// Insert activity information into the database
	_, err := db.Exec("INSERT INTO activities (id, title, proposer, startDate, endDate, maxNumber, format, description, proposeDateTime,
			acceptAdmin, acceptDateTime, applicationStatus) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", activity.activityID, activity.title, activity.proposer, 
			activity.startDate, activity.endDate, activity.maxNumber, activity.format, activity.description, activity.proposeDateTime, activity.acceptAdmin, activity.acceptDateTime, activity.applicationStatus)
	if err != nil {
		http.Error(w, "Failed to insert activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity page
	http.Redirect(w, r, "/activity", http.StatusFound)
}

func patchActivity(w http.ResponseWriter, r *http.Request) {
	// Get activity information from the request
	activity := Activity{
		activityID: r.FormValue("activityID"),
		title: r.FormValue("title"),
		proposer: r.FormValue("proposer"),
		startDate: r.FormValue("startDate"),
		endDate: r.FormValue("endDate"),
		maxNumber: r.FormValue("maxNumber"),
		format: r.FormValue("format"),
		description: r.FormValue("description"),
		proposeDateTime: r.FormValue("proposeDateTime"),
		acceptAdmin: r.FormValue("acceptAdmin"),
		acceptDateTime: r.FormValue("acceptDateTime"),
		applicationStatus: r.FormValue("applicationStatus")
	}
	
	// Update activity information in the database
	_, err := db.Exec("UPDATE activities SET title = ?, proposer = ?, startDate = ?, endDate = ?, maxNumber = ?, format = ?, description = ?, proposeDateTime = ?, 
			acceptAdmin = ?, acceptDateTime = ?, applicationStatus = ? WHERE id = ?", activity.title, activity.proposer, activity.startDate, activity.endDate, activity.maxNumber, 
			activity.format, activity.description, activity.proposeDateTime, activity.acceptAdmin, activity.acceptDateTime, activity.applicationStatus, activity.activityID)
	if err != nil {
		http.Error(w, "Failed to update activity information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity page
	http.Redirect(w, r, "/activity", http.StatusFound)
}

func deleteActivity(w http.ResponseWriter, r *http.Request) {
	// Get activity ID from the request
	activityID := r.FormValue("activityID")
	
	// Delete activity from the database
	_, err := db.Exec("DELETE FROM activities WHERE id = ?", activityID)
	if err != nil {
		http.Error(w, "Failed to delete activity: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity page
	http.Redirect(w, r, "/activity", http.StatusFound)
}

func getActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Query the database for activity roles
	var activityRoles ActivityRoles
	err := db.QueryRow("SELECT id, role FROM activity_roles WHERE id = ?", activityRoles.activityID).Scan(&activityRoles.activityID, &activityRoles.activityRole)
	if err != nil {
		http.Error(w, "Failed to fetch activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func postActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Get activity roles from the request
	activityRoles := ActivityRoles{
		activityID: r.FormValue("activityID"),
		activityRole: r.FormValue("activityRole")
	}
	
	// Insert activity roles into the database
	_, err := db.Exec("INSERT INTO activity_roles (id, role) VALUES (?, ?)", activityRoles.activityID, activityRoles.activityRole)
	if err != nil {
		http.Error(w, "Failed to insert activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity roles page
	http.Redirect(w, r, "/activityRoles", http.StatusFound)
}

func patchActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Get activity roles from the request
	activityRoles := ActivityRoles{
		activityID: r.FormValue("activityID"),
		activityRole: r.FormValue("activityRole")
	}
	
	// Update activity roles in the database
	_, err := db.Exec("UPDATE activity_roles SET role = ? WHERE id = ?", activityRoles.activityRole, activityRoles.activityID)
	if err != nil {
		http.Error(w, "Failed to update activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity roles page
	http.Redirect(w, r, "/activityRoles", http.StatusFound)
}

func deleteActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Get activity ID from the request
	activityID := r.FormValue("activityID")
	
	// Delete activity roles from the database
	_, err := db.Exec("DELETE FROM activity_roles WHERE id = ?", activityID)
	if err != nil {
		http.Error(w, "Failed to delete activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity roles page
	http.Redirect(w, r, "/activityRoles", http.StatusFound)
}

func activityRolesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getActivityRoles(w, r)
	case "POST":
		postActivityRoles(w, r)
	case "PATCH":
		patchActivityRoles(w, r)
	case "DELETE":
		deleteActivityRoles(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Query the database for activity roles
	var activityRoles ActivityRoles
	err := db.QueryRow("SELECT id, role FROM activity_roles WHERE id = ?", activityRoles.activityID).Scan(&activityRoles.activityID, &activityRoles.activityRole)
	if err != nil {
		http.Error(w, "Failed to fetch activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func postActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Get activity roles from the request
	activityRoles := ActivityRoles{
		activityID: r.FormValue("activityID"),
		activityRole: r.FormValue("activityRole")
	}
	
	// Insert activity roles into the database
	_, err := db.Exec("INSERT INTO activity_roles (id, role) VALUES (?, ?)", activityRoles.activityID, activityRoles.activityRole)
	if err != nil {
		http.Error(w, "Failed to insert activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity roles page
	http.Redirect(w, r, "/activityRoles", http.StatusFound)
}

func patchActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Get activity roles from the request
	activityRoles := ActivityRoles{
		activityID: r.FormValue("activityID"),
		activityRole: r.FormValue("activityRole")
	}
	
	// Update activity roles in the database
	_, err := db.Exec("UPDATE activity_roles SET role = ? WHERE id = ?", activityRoles.activityRole, activityRoles.activityID)
	if err != nil {
		http.Error(w, "Failed to update activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity roles page
	http.Redirect(w, r, "/activityRoles", http.StatusFound)
}

func deleteActivityRoles(w http.ResponseWriter, r *http.Request) {
	// Get activity ID from the request
	activityID := r.FormValue("activityID")
	
	// Delete activity roles from the database
	_, err := db.Exec("DELETE FROM activity_roles WHERE id = ?", activityID)
	if err != nil {
		http.Error(w, "Failed to delete activity roles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to the activity roles page
	http.Redirect(w, r, "/activityRoles", http.StatusFound)
}