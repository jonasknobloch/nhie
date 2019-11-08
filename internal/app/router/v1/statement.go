package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/internal/app"
	"github.com/neverhaveiever-io/api/internal/category"
	"github.com/neverhaveiever-io/api/internal/statement"
	"github.com/neverhaveiever-io/api/pkg/problems"
	"github.com/neverhaveiever-io/api/pkg/unique"
	"net/http"
	"strings"
)

func AddStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	var s statement.Statement

	if err := ctx.ShouldBind(&s); err != nil {
		g.ErrorResponse(problems.Default(http.StatusBadRequest))
		return
	}

	if err := s.Validate(); err != nil {
		g.ErrorResponse(problems.ValidationError(err))
		return
	}

	if err := s.Save(); err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusCreated, s)
}

func GetStatement(ctx *gin.Context) {
	if r := strings.Split(ctx.Request.URL.String(), "/random"); len(r) == 2 { // && r[1] == ""
		GetRandomStatement(ctx)
	} else {
		GetStatementByID(ctx)
	}
}

func GetStatementByID(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	// ctx.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(ctx.Params.ByName("id"))

	if err != nil {
		g.ErrorResponse(problems.Default(http.StatusBadRequest))
		return
	}

	s, err := statement.GetByID(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			g.ErrorResponse(problems.NoSuchStatement())
			return
		}

		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusOK, s)
}

func GetRandomStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	var s statement.Statement

	var categories []category.Category

	for _, v := range unique.Strings(append(ctx.QueryArray("category"), ctx.QueryArray("category[]")...)) {
		c := category.Category(v)

		if err := c.Validate(); err != nil {
			g.ErrorResponse(problems.ValidationError(err))
			return
		}

		categories = append(categories, c)
	}

	s, err := statement.GetRandomByCategory(category.GetRandom())

	if err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	// a redirect might make sense but the resulting round trip is just not worth it
	g.Response(http.StatusOK, s)
}

func DeleteStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	// ctx.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(ctx.Params.ByName("id"))

	if err != nil {
		g.ErrorResponse(problems.Default(http.StatusBadRequest))
		return
	}

	s, err := statement.GetByID(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			g.ErrorResponse(problems.NoSuchStatement())
			return
		}

		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	if err := s.Delete(); err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusNoContent)
}

func EditStatement(ctx *gin.Context) {
	g := app.Gin{C: ctx}

	// TODO: implement
	g.ErrorResponse(problems.Default(http.StatusNotImplemented))
}
