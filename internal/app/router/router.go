package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/nhie-io/api/internal/app/auth"
	v1 "github.com/nhie-io/api/internal/app/router/v1"
	"net/http"
)

func Init() error {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://nhie.io"},
	}))

	router.Route("/v1/statements", func(r chi.Router) {
		r.Get("/random", v1.GetRandomStatement)
		r.Get("/{statementID}", v1.GetStatementByID)

		r.Route("/", func(r chi.Router) {
			r.Use(middleware.BasicAuth("", auth.Accounts([]string{"admin"})))
			r.Post("/", v1.AddStatement)
			r.Put("/{statementID}", v1.EditStatement)
			r.Delete("/{statementID}", v1.DeleteStatement)
		})
	})

	return http.ListenAndServe(":8080", router)
}
