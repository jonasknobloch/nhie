package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/hostrouter"
	"net/http"
)

func Init() error {
	r := chi.NewRouter()

	hr := hostrouter.New()

	hr.Map("nhie.io", func() chi.Router {
		router := chi.NewRouter()

		router.Mount("/", webRouter())
		router.Mount("/static", staticRouter())

		return router
	}())

	hr.Map("api.nhie.io", apiRouter())

	// TODO not working with :8080

	r.Mount("/", hr)

	return http.ListenAndServe(":80", r)
}
