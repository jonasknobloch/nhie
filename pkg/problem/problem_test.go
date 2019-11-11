package problem

import (
	"net/http"
	"testing"
)

func TestDefault(t *testing.T) {
	expected := &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(http.StatusBadRequest),
		Status: http.StatusBadRequest,
	}

	if p := *Default(http.StatusBadRequest); p != *expected {
		t.Fatalf("Unexpected struct contents. %+v", p)
	}
}

func TestNoSuchStatement(t *testing.T) {
	expected := &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(http.StatusNotFound),
		Status: http.StatusNotFound,
		Detail: "Sorry, that statement does not exist.",
	}

	if p := *NoSuchStatement(); p != *expected {
		t.Fatalf("Unexpected struct contents. %+v", p)
	}
}

func TestValidationError(t *testing.T) {
	expected := &Problem{
		Type:   "about:blank",
		Title:  http.StatusText(http.StatusBadRequest),
		Status: http.StatusBadRequest,
		Errors: nil,
	}

	if p := *ValidationError(nil); p != *expected {
		t.Fatalf("Unexpected struct contents. %+v", p)
	}
}
