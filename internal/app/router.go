package app

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
)

func Init() error {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://nhie.io"},
	}))

	router.Route("/v1/statements", func(r chi.Router) {
		r.Get("/random", getRandomStatement)
		r.Get("/{statementID}", getStatementByID)

		r.Route("/", func(r chi.Router) {
			r.Use(middleware.BasicAuth("", accounts([]string{"admin"})))
			r.Post("/", addStatement)
			r.Put("/{statementID}", editStatement)
			r.Delete("/{statementID}", deleteStatement)
		})
	})

	return http.ListenAndServe(":8080", router)
}
