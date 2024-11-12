// get member function from src/server/test.go

package main

import (
	"fmt"
	"database/sql"
	"net/http"
	"html/template"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Member struct {
	memberID int(11),
	fName varchar(255),
	lName varchar(255),
	email varchar(320),
	phone varchar(20),
	githubUrl varchar(320),
	status varchar(10),
	reason text
}

var (
	db *sql.DB
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:8080)/members")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	http.HandleFunc("/member", getMember)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func getMember(w http.ResponseWriter, r *http.Request) {
    // Query the database for member information
    var member UserInfo
    err := db.QueryRow("SELECT id, name, email FROM members WHERE id = ?", userInfo.ID).Scan(&member.ID, &member.Name, &member.Email)
    if err != nil {
        http.Error(w, "Failed to fetch member information: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Load profile template
    tmpl, err := template.ParseFiles("templates/profile.html")
    if err != nil {
        http.Error(w, "Failed to load profile template: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Render template with user data
    err = tmpl.Execute(w, member)
    if err != nil {
        http.Error(w, "Failed to render profile template: "+err.Error(), http.StatusInternalServerError)
        return
    }
}

func postMember(w http.ResponseWriter, r *http.Request) {
    // Get form values
    fName := r.FormValue("fName")
    lName := r.FormValue("lName")
    email := r.FormValue("email")
    phone := r.FormValue("phone")
    githubUrl := r.FormValue("githubUrl")
    status := r.FormValue("status")
    reason := r.FormValue("reason")

    // Insert member into database
    _, err = db.Exec("INSERT INTO members (fName, lName, email, phone, githubUrl, status, reason) VALUES (?, ?, ?, ?, ?, ?, ?)", fName, lName, email, phone, githubUrl, status, reason)
    if err != nil {
        http.Error(w, "Failed to insert member: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Redirect to profile
    http.Redirect(w, r, "/profile", http.StatusFound)
}
