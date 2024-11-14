package main

import (
	"log"
	"net/http"
	"sinno-server/routes"
)

func main() {
	routes.RegisterRoutes()
	log.Println("Server started at http://localhost:" + "8080")
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
