package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Unauthorized")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if err != nil {
			log.Error(err)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
	})
}
