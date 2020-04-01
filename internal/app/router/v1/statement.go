package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/neverhaveiever-io/api/internal/app"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/statement"
	"github.com/neverhaveiever-io/api/internal/translate"
	"github.com/neverhaveiever-io/api/pkg/problem"
	"github.com/neverhaveiever-io/api/pkg/unique"
	"net/http"
	"strings"
)

func AddStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	var s statement.Statement

	if err := ctx.ShouldBind(&s); err != nil {
		g.ErrorResponse(problem.Default(http.StatusBadRequest))
		return
	}

	if err := s.Validate(); err != nil {
		g.ErrorResponse(problem.ValidationError(err))
		return
	}

	if err := s.Save(); err != nil {

		_ = g.C.Error(err)

		// catch unique_violation with error code 23505
		if e, ok := err.(*pq.Error); ok && e.Code == "23505" {
			g.ErrorResponse(problem.StatementAlreadyExists())
			return
		}

		g.ErrorResponse(problem.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusCreated, s)
}

func GetStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	var s *statement.Statement
	var err error

	if r := strings.Split(ctx.Request.URL.String(), "/random"); len(r) == 2 { // && r[1] == ""
		s, err = getRandomStatement(g)
	} else {
		s, err = getStatementByID(g)
	}

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			g.ErrorResponse(problem.NoSuchStatement())
			return
		}

		_ = g.C.Error(err)
		g.ErrorResponse(problem.Default(http.StatusInternalServerError))
		return
	}

	// error already handled
	if s == nil {
		return
	}

	if g.C.Query("language") != "" {
		tags := unique.Strings(append(ctx.QueryArray("language"), ctx.QueryArray("language[]")...))
		s.FetchTranslations(translate.MatchTags(tags...)...)
	}

	// a redirect might make sense but the resulting round trip is just not worth it
	g.Response(http.StatusOK, s)
}

func getStatementByID(g app.Gin) (*statement.Statement, error) {
	// g.C.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(g.C.Params.ByName("id"))

	if err != nil {
		g.ErrorResponse(problem.Default(http.StatusBadRequest))
		return nil, nil
	}

	return statement.GetByID(id)
}

func getRandomStatement(g app.Gin) (*statement.Statement, error) {
	var categories []category.Category

	for _, v := range unique.Strings(append(g.C.QueryArray("category"), g.C.QueryArray("category[]")...)) {
		c := category.Category(v)

		if err := c.Validate(); err != nil {
			g.ErrorResponse(problem.ValidationError(err))
			return nil, nil
		}

		categories = append(categories, c)
	}

	return statement.GetRandomByCategory(category.GetRandom(categories...))
}

func DeleteStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	// ctx.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(ctx.Params.ByName("id"))

	if err != nil {
		g.ErrorResponse(problem.Default(http.StatusBadRequest))
		return
	}

	s, err := statement.GetByID(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			g.ErrorResponse(problem.NoSuchStatement())
			return
		}

		_ = g.C.Error(err)
		g.ErrorResponse(problem.Default(http.StatusInternalServerError))
		return
	}

	if err := s.Delete(); err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problem.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusNoContent)
}

func EditStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	// TODO: implement
	g.ErrorResponse(problem.Default(http.StatusNotImplemented))
}
