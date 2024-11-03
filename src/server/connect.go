package main
//import essentials libraries
import (
	"fmt",
	"database/sql",
	"net/http",
	"html/template",
	"log",

	_ "github.com/go-sql-driver/mysql"
}

var (
	db *sql.DB
)

func main() {
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