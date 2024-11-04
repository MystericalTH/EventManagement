package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Workshop struct to hold workshop details
type Workshop struct {
	WorkshopID int
	StartTime  time.Time
	EndTime    time.Time
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

	// Display fetched workshop data
	fmt.Fprintf(w, "Workshop: %+v", workshop)
}

func postWorkshop(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form data: "+err.Error(), http.StatusBadRequest)
		return
	}

	startTime, err := time.Parse("2006-01-02 15:04:05", r.FormValue("startTime"))
	if err != nil {
		http.Error(w, "Invalid startTime format", http.StatusBadRequest)
		return
	}

	endTime, err := time.Parse("2006-01-02 15:04:05", r.FormValue("endTime"))
	if err != nil {
		http.Error(w, "Invalid endTime format", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("INSERT INTO workshop (startTime, endTime) VALUES (?, ?)", startTime, endTime)
	if err != nil {
		http.Error(w, "Failed to insert workshop information: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Workshop created successfully"))
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
