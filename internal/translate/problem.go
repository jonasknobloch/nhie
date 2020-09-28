package translate

import (
	"fmt"
	"github.com/nhie-io/api/pkg/problem"
	"golang.org/x/text/language/display"
	"net/http"
)

func NewMatchingErrorProblem(e *MatchingError) *problem.Problem {
	p := problem.Default(http.StatusBadRequest)

	switch e.err {
	case ErrInvalidLanguageTag:
		p.Detail = fmt.Sprintf("\"%s\" is not a valid language tag.", e.input)
	case ErrUnknownLanguageTag:
		p.Detail = fmt.Sprintf("The received language tag \"%s\" is not known to us.", e.input)
	case ErrUnsupportedLanguage:
		p.Detail = fmt.Sprintf("%s is currently not supported.", display.English.Tags().Name(e.tag))
	}

	return p
}
