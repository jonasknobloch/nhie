package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/jonasknobloch/nhie/web"
	"net/http"
)

func staticRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/*", http.StripPrefix("/static/", http.FileServer(http.FS(web.Build))).ServeHTTP)

	return router
}
