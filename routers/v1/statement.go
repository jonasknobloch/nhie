package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/neverhaveiever-io/api/models"
	"github.com/neverhaveiever-io/api/pkg/app"
	"github.com/neverhaveiever-io/api/pkg/problems"
	"github.com/neverhaveiever-io/api/pkg/unique"
	"net/http"
	"strings"
)

func AddStatement(c *gin.Context) {
	g := app.Gin{C: c}

	var statement models.Statement

	if err := c.ShouldBind(&statement); err != nil {
		g.ErrorResponse(problems.Default(http.StatusBadRequest))
		return
	}

	if err := statement.Validate(); err != nil {
		g.ErrorResponse(problems.ValidationError(err))
		return
	}

	if err := statement.Save(); err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusCreated, statement)
}

func GetStatement(c *gin.Context) {
	if r := strings.Split(c.Request.URL.String(), "/random"); len(r) == 2 { // && r[1] == ""
		GetRandomStatement(c)
	} else {
		GetStatementById(c)
	}
}

func GetStatementById(c *gin.Context) {
	g := app.Gin{C: c}

	// c.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(c.Params.ByName("id"))

	if err != nil {
		g.ErrorResponse(problems.Default(http.StatusBadRequest))
		return
	}

	statement, err := models.GetStatementById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			g.ErrorResponse(problems.NoSuchStatement())
			return
		}

		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusOK, statement)
}

func GetRandomStatement(c *gin.Context) {
	g := app.Gin{C: c}

	var statement models.Statement

	var categories []models.Category

	for _, v := range unique.Strings(append(c.QueryArray("category"), c.QueryArray("category[]")...)) {
		category := models.Category(v)

		if err := category.Validate(); err != nil {
			g.ErrorResponse(problems.ValidationError(err))
			return
		}

		categories = append(categories, category)
	}

	statement, err := models.GetRandomStatementByCategory(models.GetRandomCategory())

	if err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	// a redirect might make sense but the resulting round trip is just not worth it
	g.Response(http.StatusOK, statement)
}

func DeleteStatement(c *gin.Context) {
	g := app.Gin{C: c}

	// c.Params.ByName("id") returns an empty string if no matching key is found
	id, err := uuid.Parse(c.Params.ByName("id"))

	if err != nil {
		g.ErrorResponse(problems.Default(http.StatusBadRequest))
		return
	}

	statement, err := models.GetStatementById(id)

	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			g.ErrorResponse(problems.NoSuchStatement())
			return
		}

		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	if err := statement.Delete(); err != nil {
		_ = g.C.Error(err)
		g.ErrorResponse(problems.Default(http.StatusInternalServerError))
		return
	}

	g.Response(http.StatusNoContent)
}

func EditStatement(c *gin.Context) {
	g := app.Gin{C: c}

	// TODO: implement
	g.ErrorResponse(problems.Default(http.StatusNotImplemented))
}
