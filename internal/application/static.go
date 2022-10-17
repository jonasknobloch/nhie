package application

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func staticRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))).ServeHTTP)

	return router
}
