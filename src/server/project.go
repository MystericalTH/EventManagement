package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Project struct {
	ProjectID int    `json:"projectID"`
	Advisor   string `json:"advisor"`
}

// Global variable for the database connection

func project() {
	// Initialize the database connection (replace with your DSN)
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	http.HandleFunc("/project", projectHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

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
	// Extract project ID from the URL query parameters
	projectID := r.URL.Query().Get("projectID")
	if projectID == "" {
		http.Error(w, "projectID not provided", http.StatusBadRequest)
		return
	}

	// Query the database for project information
	var project Project
	err := db.QueryRow("SELECT projectID, advisor FROM project WHERE projectID = ?", projectID).Scan(&project.ProjectID, &project.Advisor)
	if err != nil {
		http.Error(w, "Failed to fetch project information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return project information as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(project)
}

func postProject(w http.ResponseWriter, r *http.Request) {
	// Decode the JSON request body into a Project struct
	var project Project
	if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Check that the advisor field is provided
	if project.Advisor == "" {
		http.Error(w, "Advisor field is required", http.StatusBadRequest)
		return
	}

	// Insert project information into the database
	result, err := db.Exec("INSERT INTO project (advisor) VALUES (?)", project.Advisor)
	if err != nil {
		http.Error(w, "Failed to insert project information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the inserted project's ID
	projectID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, "Failed to retrieve last insert ID: "+err.Error(), http.StatusInternalServerError)
		return
	}
	project.ProjectID = int(projectID)

	// Respond with the created project as JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}
