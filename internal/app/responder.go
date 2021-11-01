package app

import (
	"bytes"
	"encoding/json"
	"github.com/go-chi/render"
	"github.com/nhie-io/api/pkg/problem"
	"net/http"
)

func Responder(w http.ResponseWriter, r *http.Request, v interface{}) {
	if _, ok := v.(*problem.Problem); ok {
		// TODO log error -> add error instance to problem

		Problem(w, r, v)
		return
	}

	render.JSON(w, r, v)
}

func Problem(w http.ResponseWriter, r *http.Request, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)

	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", problem.MediaType+"; charset=utf-8")

	if status, ok := r.Context().Value(render.StatusCtxKey).(int); ok {
		w.WriteHeader(status)
	}

	w.Write(buf.Bytes())
}
