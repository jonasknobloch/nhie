package problem

import (
	"net/http"
)

const MediaType = "application/problem+json"

type Problem struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status,omitempty"`
	Detail   string `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
	Errors   error  `json:"validation-errors,omitempty"`
}

func Default(status int) *Problem {
	return &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(status),
		Status: status,
	}
}

func NoSuchStatement() *Problem {
	return &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(http.StatusNotFound),
		Status: http.StatusNotFound,
		Detail: "Sorry, that statement does not exist.",
	}
}

func ValidationError(errors error) *Problem {
	return &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(http.StatusBadRequest),
		Status: http.StatusBadRequest,
		Errors: errors,
	}
}

func StatementAlreadyExists() *Problem {
	return &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(http.StatusConflict),
		Status: http.StatusConflict,
		Detail: "Sorry, that statement already exists.",
	}
}
