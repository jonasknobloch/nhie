package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"net/http"
)

func apiRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/v1/statements/random", func(w http.ResponseWriter, r *http.Request) {
		c, ok := queryCategories(r)

		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		s, err := statement.GetRandomByCategory(c.Random())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		l, ok := queryLanguage(r)

		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if l != translate.SourceLanguage {
			err := s.Translate(l)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		render.JSON(w, r, s)
	})

	return router
}
