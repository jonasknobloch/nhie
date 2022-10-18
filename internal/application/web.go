package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/database"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

var templates *template.Template

func init() {
	templates = template.Must(template.ParseFiles("web/index.html"))
}

func webRouter() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		c := category.NewSelection()

		c.Add(category.Harmless)
		c.Add(category.Delicate)
		c.Add(category.Offensive)

		s, err := statement.GetRandomByCategory(c.Random())

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		renderGame(s, w, r)
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

		s, err := statement.GetByID(uuid.MustParse(rID))

		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		renderGame(s, w, r)
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
		statement:  s,
		categories: c,
		language:   l,
	}

	w.Header().Add("Content-Location", data.ContentLocation())

	if err := templates.ExecuteTemplate(w, "index.html", data); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
