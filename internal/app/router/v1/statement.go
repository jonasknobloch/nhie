package v1

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/nhie-io/api/internal/app"
	"github.com/nhie-io/api/internal/cache"
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/history"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"github.com/nhie-io/api/pkg/problem"
	"github.com/nhie-io/api/pkg/unique"
	"gorm.io/gorm"
	"net/http"
)

func init() {
	render.Respond = app.Responder
}

func AddStatement(w http.ResponseWriter, r *http.Request) {
	var s statement.Statement

	if err := render.Bind(r, &s); err != nil {
		Render(w, r, problem.Default(http.StatusBadRequest))
		return
	}

	if err := s.Validate(); err != nil {
		Render(w, r, problem.ValidationError(err))
		return
	}

	if err := s.Save(); err != nil {

		// catch unique_violation with error code 23505
		if e, ok := err.(*pgconn.PgError); ok && e.Code == "23505" {
			Render(w, r, problem.StatementAlreadyExists())
			return
		}

		Render(w, r, problem.Default(http.StatusInternalServerError))
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, &s)
}

func GetStatementByID(w http.ResponseWriter, r *http.Request) {
	// g.C.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(chi.URLParam(r, "statementID"))

	if err != nil {
		Render(w, r, problem.Default(http.StatusBadRequest))
		return
	}

	s, err := statement.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Render(w, r, problem.NoSuchStatement())
		}

		Render(w, r, problem.Default(http.StatusInternalServerError))
		return
	}

	if l := r.URL.Query().Get("language"); l != "" {
		if p := TranslateStatement(s, l); p != nil {
			Render(w, r, p)
			return
		}
	}

	render.JSON(w, r, s)
}

func GetRandomStatement(w http.ResponseWriter, r *http.Request) {
	var categories []category.Category

	q := r.URL.Query()
	c := make([]string, 0)

	if v, ok := q["category"]; ok {
		c = append(c, v...)
	}

	if v, ok := q["category[]"]; ok {
		c = append(c, v...)
	}

	for _, v := range unique.Strings(c) {
		c := category.Category(v)

		if err := c.Validate(); err != nil {
			Render(w, r, problem.ValidationError(err))
			return
		}

		categories = append(categories, c)
	}

	s, p, err := statement.GetRandomByCategory(category.GetRandom(categories...))

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Render(w, r, problem.NoSuchStatement())
		}

		Render(w, r, problem.Default(http.StatusInternalServerError))
		return
	}

	if r.URL.Query().Get("game_id") != "" {
		var gameID uuid.UUID
		gameID, err = uuid.Parse(r.URL.Query().Get("game_id"))

		if err != nil {
			Render(w, r, problem.Default(http.StatusBadRequest))
			return
		}

		var e bool
		maxTries := history.MaxTries

		for try := 0; try < maxTries; try++ {

			if e, err = history.Exists(gameID, s); err != nil {
				break
			}

			if err = history.Add(gameID, s); err != nil {
				break
			}

			if !e {
				break
			}

			history.ReportDuplicate(try+1, maxTries, p)

			s, p, err = statement.GetRandomByCategory(category.GetRandom(categories...))

			if err != nil {
				break
			}
		}
	}

	if l := r.URL.Query().Get("language"); l != "" {
		if p := TranslateStatement(s, l); p != nil {
			Render(w, r, p)
			return
		}
	}

	render.JSON(w, r, s)
}

func TranslateStatement(s *statement.Statement, l string) *problem.Problem {
	matchedTag, err := translate.MatchTag(l)

	if err != nil {
		var e *translate.MatchingError
		if errors.As(err, &e) {
			return translate.NewMatchingErrorProblem(e)
		} else {
			return problem.Default(http.StatusInternalServerError)
		}
	}

	if matchedTag != translate.SourceLanguage {
		if err := s.Translate(matchedTag); err != nil {
			// might be just cache error
			var e *cache.Error
			if !errors.As(err, &e) {
				return problem.Default(http.StatusInternalServerError)
			}
		}
	}

	return nil
}

func DeleteStatement(w http.ResponseWriter, r *http.Request) {
	// ctx.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(chi.URLParam(r, "statementID"))

	if err != nil {
		Render(w, r, problem.Default(http.StatusBadRequest))
		return
	}

	s, err := statement.GetByID(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			Render(w, r, problem.NoSuchStatement())
			return
		}

		Render(w, r, problem.Default(http.StatusInternalServerError))
		return
	}

	if err := s.Delete(); err != nil {
		Render(w, r, problem.Default(http.StatusInternalServerError))
		return
	}

	render.NoContent(w, r)
}

func EditStatement(w http.ResponseWriter, r *http.Request) {
	// TODO: implement

	Render(w, r, problem.Default(http.StatusNotImplemented))
}
