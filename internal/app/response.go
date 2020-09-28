package app

import (
	"github.com/nhie-io/api/pkg/problem"
)

func (g *Gin) Response(httpCode int, data ...interface{}) {
	if data == nil {
		g.C.JSON(httpCode, data)
	} else {
		g.C.JSON(httpCode, data[0])
	}
}

func (g *Gin) ErrorResponse(p *problem.Problem) {
	g.C.Header("Content-Type", problem.MediaType+"; charset=utf-8")
	g.C.JSON(p.Status, p)
}
