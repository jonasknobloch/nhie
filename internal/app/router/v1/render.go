package v1

import (
	"github.com/go-chi/render"
	"net/http"
)

func Render(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	if err := render.Render(w, r, v); err != nil {
		// TODO error log
	}
}
