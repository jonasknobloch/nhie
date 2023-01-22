package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jonasknobloch/nhie/internal/statement"
	"github.com/jonasknobloch/nhie/internal/translate"
	"net/http"
)

func apiRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{URLScheme + "://" + WebHost},
	}))

	router.Get("/v1/statements/random", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/v2/statements/next?"+r.URL.RawQuery, http.StatusPermanentRedirect)
	})

	router.Get("/v2/statements/next", func(w http.ResponseWriter, r *http.Request) {
		c, ok := queryCategories(r)

		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		sID, ok := queryStatementID(r)

		var s *statement.Statement
		var err error

		if !ok {
			s, err = statement.GetRandomByCategory(c.Random())
		} else {
			s, err = statement.GetNextByPreviousIDAndCategory(sID, c.Random())
		}

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
