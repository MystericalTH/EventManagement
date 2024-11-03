package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"

	_ "github.com/go-sql-driver/mysql"
)

type chatDevAd struct {
	messageID int(11),
	adminID int(11),
	developerID int(11),
	message text,
	datetime datetime
}

type Developer struct {
	developerID int(11),
	email varchar(320)
}

type Admin struct {
	adminID int(11),
	email varchar(320)
}

var (
	db *sql.DB
)

func chatHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getChat(w, r)
	case "POST":
		postChat(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getChat(w http.ResponseWriter, r *http.Request) {
	// Query the database for chat information
	var chat chatDevAd
	err := db.QueryRow("SELECT id, adminID, developerID, message, datetime FROM chat WHERE id = ?", chat.messageID).Scan(&chat.messageID, &chat.adminID, &chat.developerID, &chat.message, &chat.datetime)
	if err != nil {
		http.Error(w, "Failed to fetch chat information: "+err.Error(), http.StatusInternalServerError)
		return
}

func postChat(w http.ResponseWriter, r *http.Request) {
	// Insert chat information into the database
	var chat chatDevAd
	err := db.QueryRow("INSERT INTO chat (adminID, developerID, message, datetime) VALUES (?, ?, ?, ?)", chat.adminID, chat.developerID, chat.message, chat.datetime)
	if err != nil {
		http.Error(w, "Failed to insert chat information: "+err.Error(), http.StatusInternalServerError)
		return
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getAdmin(w, r)
	case "POST":
		postAdmin(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getAdmin(w http.ResponseWriter, r *http.Request) {
	// Query the database for admin information
	var admin Admin
	err := db.QueryRow("SELECT id, email FROM admin WHERE id = ?", admin.adminID).Scan(&admin.adminID, &admin.email)
	if err != nil {
		http.Error(w, "Failed to fetch admin information: "+err.Error(), http.StatusInternalServerError)
		return
}

func postAdmin(w http.ResponseWriter, r *http.Request) {
	// Insert admin information into the database
	var admin Admin
	err := db.QueryRow("INSERT INTO admin (email) VALUES (?)", admin.email)
	if err != nil {
		http.Error(w, "Failed to insert admin information: "+err.Error(), http.StatusInternalServerError)
		return
}

func developerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getDeveloper(w, r)
	case "POST":
		postDeveloper(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getDeveloper(w http.ResponseWriter, r *http.Request) {
	// Query the database for developer information
	var developer Developer
	err := db.QueryRow("SELECT id, email FROM developer WHERE id = ?", developer.developerID).Scan(&developer.developerID, &developer.email)
	if err != nil {
		http.Error(w, "Failed to fetch developer information: "+err.Error(), http.StatusInternalServerError)
		return
}

func postDeveloper(w http.ResponseWriter, r *http.Request) {
	// Insert developer information into the database
	var developer Developer
	err := db.QueryRow("INSERT INTO developer (email) VALUES (?)", developer.email)
	if err != nil {
		http.Error(w, "Failed to insert developer information: "+err.Error(), http.StatusInternalServerError)
		return
}