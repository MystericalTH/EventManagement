package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Server starting on :3306...")

	err := http.ListenAndServe("localhost:3306", r)
	if err != nil {
		log.Error(err)
	}

}
