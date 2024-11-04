package main

//import essentials libraries
import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func connect() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/activitymanagement")
	if err != nil {
		panic(err.Error())
	}
	// the upper code is to connect to the database and error handling
	// Ensure the connection is valid
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	http.HandleFunc("/member", memberHandler)
	http.HandleFunc("/activity", activityHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
