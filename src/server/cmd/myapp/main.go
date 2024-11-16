package main

import (
	"database/sql"
	"fmt"
	"log"
	"sinno-server/pkg/api" // Import the api package
	"sinno-server/pkg/db"  // Import your db package

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	// Initialize the MySQL database connection
	conn, err := sql.Open("mysql", "root:Palioandy38550@tcp(127.0.0.1:3306)/myActManage")
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}

	// Ping the database to check the connection
	if err := conn.Ping(); err != nil {
		log.Fatal("Error pinging database: ", err)
	}
	fmt.Println("Database connection successful!")

	// Create a new instance of db.Queries using the connection
	queries := db.New(conn)

	// Initialize your Gin router
	r := gin.Default()

	// Register routes using RegisterRoutes function
	api.RegisterRoutes(r, queries)

	// Run the server
	r.Run(":8080")
}
