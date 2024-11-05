package api

import (
	"encoding/json"
	"net/http"
)

type Member struct {
	MemberID int    `json:"memberID"`
	FName    string `json:"fName"`
}

type Error struct {
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, message string, code int) {

	resp := Error{Message: message}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {

		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "Internal Server Error", http.StatusInternalServerError)

	}
)
