package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Init() {
	var err error
	// user := "root"
	// password := "root"
	// host := "localhost"
	// port := "3306"
	// dbname := "ClubManagement"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	// 	// os.Getenv("DB_USER"),
	// 	// os.Getenv("DB_PASSWORD"),
	// 	// os.Getenv("DB_HOST"),
	// 	// os.Getenv("DB_PORT"),
	// 	// os.Getenv("DB_NAME"),
	// 	user,
	// 	password,
	// 	host,
	// 	port,
	// 	dbname,
	// )

	dsn := "root:root@tcp(mysql:3306)/ClubManagement"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
}
