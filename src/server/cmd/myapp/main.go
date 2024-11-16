package main

import (
	"database/sql"
	"fmt"
	"log"
	"sinno-server/pkg/api/handler"
	"sinno-server/pkg/db" // Import your db package

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // MySQL driver
)

func main() {
	// Initialize the MySQL database connection
	conn, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/mydb")
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

	// Pass the queries to the handlers
	r.GET("/members", func(c *gin.Context) {
		handler.GetAllMembers(c, queries)
	})
	r.GET("/members/:id", func(c *gin.Context) {
		handler.GetMemberByID(c, queries)
	})
	r.POST("/members", func(c *gin.Context) {
		handler.CreateMember(c, queries)
	})
	r.PUT("/members/:id/accept", func(c *gin.Context) {
		handler.AcceptMember(c, queries)
	})

	// Run the server
	r.Run(":8080")
}
