package main

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"sinno-server/routes"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	// Replace with your actual Google credentials
	port        = os.Getenv("LISTEN_PORT")
	oauthConfig = &oauth2.Config{
		RedirectURL:  fmt.Sprintf("http://localhost:%s/auth/google/callback", port),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),     // Set in your environment
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"), // Set in your environment
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint:     google.Endpoint,
	}
	sessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
	sessionName  = "session-one"
)

var (
	db *sql.DB
)

type UserInfo struct {
	Name  string
	Email string
}

func init() {
	// Register UserInfo type to store in session
	gob.Register(UserInfo{})
}

func main() {
	routes.RegisterRoutes()
	log.Println("Server started at http://localhost:" + "8080")
	log.Fatal(http.ListenAndServe(":"+"8080", nil))
}
