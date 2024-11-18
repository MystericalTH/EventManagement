package main

import (
	"log"
	"net/http"
	"os"

	"sinno-server/pkg/api"
	"sinno-server/pkg/db"

	"github.com/gorilla/mux"
)

var (
	port = os.Getenv("LISTEN_PORT")
)

func main() {
	db.Init()
	api.ActivitiesRoutes(mux.NewRouter())
	api.LogRoutes()
	log.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
