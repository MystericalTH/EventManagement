package routes

import (
	"net/http"
	"sinno-server/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/api/login", handlers.HandleLogin)
	http.HandleFunc("/api/auth/google/callback", handlers.HandleCallback)
	http.HandleFunc("/api/logout", handlers.HandleLogout)
	http.HandleFunc("/api/verify", handlers.HandleVerifyRole)
}
