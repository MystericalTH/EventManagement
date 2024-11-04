package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Project struct {
	ProjectID int
	Advisor   string
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

	// Render project information as needed, e.g., to a template or JSON (assuming JSON here for simplicity)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"projectID": %d, "advisor": "%s"}`, project.ProjectID, project.Advisor)
}

func postProject(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Retrieve advisor value from form
	advisor := r.FormValue("advisor")
	if advisor == "" {
		http.Error(w, "Advisor field is required", http.StatusBadRequest)
		return
	}

	// Insert project information into the database
	_, err := db.Exec("INSERT INTO project (advisor) VALUES (?)", advisor)
	if err != nil {
		http.Error(w, "Failed to insert project information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Project created successfully"))
}

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
