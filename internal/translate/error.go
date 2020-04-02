package translate

import (
	"errors"
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var (
	ErrInvalidLanguageTag    = errors.New("invalid language tag")
	ErrUnknownLanguageTag    = errors.New("unknown language tag")
	ErrUnsupportedLanguage   = errors.New("unsupported language")
	ErrNoTranslationReceived = errors.New("no translation received")
)

type Error struct {
	error
}

func newError(err error) *Error {
	return &Error{err}
}

func (e *Error) Error() string {
	return "translate -> " + e.error.Error()
}

func (e *Error) Unwrap() error {
	return e.error
}

type MatchingError struct {
	err   error
	input string
	tag   language.Tag
}

func newMatchingError(err error, input string, tag language.Tag) *MatchingError {
	return &MatchingError{
		err:   err,
		input: input,
		tag:   tag,
	}
}

func (e *MatchingError) Error() string {
	if e.err == ErrUnsupportedLanguage {
		return fmt.Sprintf("matching -> %s: %s [%s]", e.err.Error(), display.English.Tags().Name(e.tag), e.tag.String())
	}

	return fmt.Sprintf("matching -> %s: %s", e.err.Error(), e.input)
}

func (e *MatchingError) Unwrap() error {
	return e.err
}
