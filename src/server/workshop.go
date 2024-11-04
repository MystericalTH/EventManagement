package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Workshop struct to hold workshop details
type Workshop struct {
	WorkshopID int       `json:"workshopID"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

func workshop() {
	// Initialize the database connection (replace with your DSN)
	var err error
	db, err = sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	http.HandleFunc("/workshop", workshopHandler)
	log.Println("Server starting on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func workshopHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getWorkshop(w, r)
	case "POST":
		postWorkshop(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getWorkshop(w http.ResponseWriter, r *http.Request) {
	workshopID := r.URL.Query().Get("workshopID")
	if workshopID == "" {
		http.Error(w, "workshopID not provided", http.StatusBadRequest)
		return
	}

	var workshop Workshop
	err := db.QueryRow("SELECT workshopID, startTime, endTime FROM workshop WHERE workshopID = ?", workshopID).
		Scan(&workshop.WorkshopID, &workshop.StartTime, &workshop.EndTime)
	if err != nil {
		http.Error(w, "Failed to fetch workshop information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return workshop data as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(workshop)
}

func postWorkshop(w http.ResponseWriter, r *http.Request) {
	var workshop Workshop
	if err := json.NewDecoder(r.Body).Decode(&workshop); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Validate and parse time fields
	if workshop.StartTime.IsZero() || workshop.EndTime.IsZero() {
		http.Error(w, "StartTime and EndTime must be provided and valid", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO workshop (startTime, endTime) VALUES (?, ?)", workshop.StartTime, workshop.EndTime)
	if err != nil {
		http.Error(w, "Failed to insert workshop information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Workshop created successfully"))
}
