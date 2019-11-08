package app

import (
	"github.com/neverhaveiever-io/api/pkg/problems"
)

func (g *Gin) Response(httpCode int, data ...interface{}) {
	if data == nil {
		g.C.JSON(httpCode, data)
	} else {
		g.C.JSON(httpCode, data[0])
	}
}

func (g *Gin) ErrorResponse(problem *problems.Problem) {
	g.C.Header("Content-Type", problems.MediaType+"; charset=utf-8")
	g.C.JSON(problem.Status, problem)
}
