package problem

import (
	"github.com/go-chi/render"
	"net/http"
)

func (p *Problem) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, p.Status)

	return nil
}
