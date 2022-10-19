package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/hostrouter"
	"net/http"
)

func Init(webHost, apiHost string) error {
	r := chi.NewRouter()

	r.Use(middleware.Compress(5, "text/html"))
	r.Use(middleware.Heartbeat("/ping"))

	hr := hostrouter.New()

	hr.Map(webHost, func() chi.Router {
		router := chi.NewRouter()

		router.Mount("/", webRouter())
		router.Mount("/static", staticRouter())

		return router
	}())

	hr.Map(apiHost, apiRouter())

	// TODO not working with :8080

	r.Mount("/", hr)

	return http.ListenAndServe(":80", r)
}
