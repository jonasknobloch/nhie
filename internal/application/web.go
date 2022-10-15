package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/database"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"net/url"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("web/index.html"))
}

func webRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))).ServeHTTP)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		c := category.NewSelection()

		c.Add(category.Harmless)
		c.Add(category.Delicate)
		c.Add(category.Offensive)

		s, err := statement.GetRandomByCategory(c.Random())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		u, err := url.Parse("/statements/" + s.ID.String())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		http.Redirect(w, r, u.String(), http.StatusTemporaryRedirect)
	})

	router.Get("/statements/next", func(w http.ResponseWriter, r *http.Request) {
		sID, ok := queryStatementID(r)

		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		c, ok := queryCategories(r)

		if !ok {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var pos int
		var rID string

		if err := database.C.Transaction(func(tx *gorm.DB) error {
			if err := tx.Raw(`SELECT position FROM game WHERE id = ?;`, sID).Scan(&pos).Error; err != nil {
				return err
			}

			if err := tx.Raw(`SELECT id
						FROM (SELECT * FROM game WHERE position > ? UNION ALL SELECT * FROM game WHERE position < ?) AS game
						WHERE category = ?
						LIMIT 1;`, pos, pos, c.Random()).Scan(&rID).Error; err != nil {
				return err
			}

			return nil
		}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		var redirect string

		{
			redirectURL := *r.URL

			redirectURL.Path = "/statements/" + rID

			if q := redirectURL.Query(); q.Has("statement_id") {
				q.Del("statement_id")
				redirectURL.RawQuery = q.Encode()
			}

			redirect = redirectURL.RequestURI()
		}

		http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
	})

	router.Get("/statements/{statementID}", func(w http.ResponseWriter, r *http.Request) {
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

		if l != translate.SourceLanguage {
			err := s.Translate(l)

			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}

		if err := templates.ExecuteTemplate(w, "index.html", state{
			statement:  s,
			categories: c,
			language:   l,
		}); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	})

	return router
}
