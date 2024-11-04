package main

import (
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	_ "github.com/go-sql-driver/mysql"

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
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/auth/google/callback", handleCallback)

	http.HandleFunc("/profile", handleProfile)
	http.HandleFunc("/logout", handleLogout)

	log.Println("Server started at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "Welcome! <a href='/api/login'>Login with Google</a>")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	// Redirect user to Google's OAuth consent page
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	// Retrieve the authorization code from Google

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Authorization code not found", http.StatusBadRequest)
		return
	}

	// Exchange the code for a token
	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Use the token to fetch user info from Google
	client := oauthConfig.Client(context.Background(), token)
	userInfoResp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		http.Error(w, "Failed tm get user info: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	// Parse the user info
	var userInfo UserInfo
	if err := json.NewDecoder(userInfoResp.Body).Decode(&userInfo); err != nil {
		http.Error(w, "Failed to decode user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Store user info in session
	session, _ := sessionStore.Get(r, sessionName)
	session.Values["user"] = userInfo
	session.Save(r, w)

	// Redirect to profile
	http.Redirect(w, r, "/profile", http.StatusFound)
}

func handleProfile(w http.ResponseWriter, r *http.Request) {
	// Get the session
	session, _ := sessionStore.Get(r, sessionName)

	// Check if user is authenticated
	userInfo, ok := session.Values["user"].(UserInfo)
	if !ok {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	// Load profile template
	tmpl, err := template.ParseFiles("templates/profile.html")
	if err != nil {
		http.Error(w, "Failed to load profile template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Render template with user data
	err = tmpl.Execute(w, userInfo)
	if err != nil {
		http.Error(w, "Failed to render profile template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
	// Clear the session

	session, _ := sessionStore.Get(r, sessionName)
	session.Options.MaxAge = -1
	session.Save(r, w)

	// Redirect to home
	http.Redirect(w, r, "/", http.StatusFound)
}
