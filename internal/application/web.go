package application

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/jonasknobloch/nhie/internal/category"
	"github.com/jonasknobloch/nhie/internal/statement"
	"github.com/jonasknobloch/nhie/internal/translate"
	"github.com/jonasknobloch/nhie/web"
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.New("index").Parse(web.Index))
}

func webRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(NegotiateLanguage)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		c := category.NewSelection()

		c.Add(category.Harmless)
		c.Add(category.Delicate)
		c.Add(category.Offensive)

		s, err := statement.GetRandomByCategory(c.Random())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		http.Redirect(w, r, fmt.Sprintf("/statements/%s", s.ID.String()), http.StatusSeeOther)
	})

	router.Get("/statements/{statementID}", func(w http.ResponseWriter, r *http.Request) {
		sID, err := uuid.Parse(chi.URLParam(r, "statementID"))

		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		s, err := statement.GetByID(sID)

		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		renderGame(s, w, r)
	})

	return router
}

func renderGame(s *statement.Statement, w http.ResponseWriter, r *http.Request) {
	l, ok := queryLanguage(r)

	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	c, ok := queryCategories(r)

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

	data := state{
		statement:   s,
		categories:  c,
		language:    l,
		preferences: &preferences{},
		config:      &config{},
	}

	data.preferences.InvertColorScheme, _ = queryInvertColorScheme(r)

	data.config.WebHost = WebHost
	data.config.ApiHost = ApiHost
	data.config.URLScheme = URLScheme

	w.Header().Set("Content-Type", "text/html")

	if err := templates.ExecuteTemplate(w, "index", data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
