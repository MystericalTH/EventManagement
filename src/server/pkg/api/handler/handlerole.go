package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var (
	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionName  = "session-one"
)

func HandleVerifyRole(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, sessionName)
	role, ok := session.Values["role"].(string)
	if !ok {
		role = "unknown"
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"role\":\"%s\"}", role)
}
