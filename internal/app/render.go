package app

import (
	"github.com/go-chi/render"
	"net/http"
)

func renderJSON(w http.ResponseWriter, r *http.Request, v render.Renderer) {
	if err := render.Render(w, r, v); err != nil {
		// TODO error log
	}
}
