package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Project struct {
	projectID int(11),
	advisor varchar(255)
}

var (
	db *sql.DB
)

func projectHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProject(w, r)
	case "POST":
		postProject(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProject(w http.ResponseWriter, r *http.Request) {
	// Query the database for project information
	var project Project
	err := db.QueryRow("SELECT id, advisor FROM project WHERE id = ?", project.projectID).Scan(&project.projectID, &project.advisor)
	if err != nil {
		http.Error(w, "Failed to fetch project information: "+err.Error(), http.StatusInternalServerError)
		return
}

func postProject(w http.ResponseWriter, r *http.Request) {
	// Insert project information into the database
	var project Project
	err := db.QueryRow("INSERT INTO project (advisor) VALUES (?)", project.advisor)
	if err != nil {
		http.Error(w, "Failed to insert project information: "+err.Error(), http.StatusInternalServerError)
		return
	}
}