package handlers

import (
	"/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/member", func(router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/fName", memberHandler)
	})
}
