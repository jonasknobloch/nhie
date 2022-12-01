package application

import (
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"golang.org/x/text/language"
)

type state struct {
	statement   *statement.Statement
	categories  *category.Selection
	language    language.Tag
	preferences *preferences
}

type categories struct {
	Harmless  string
	Delicate  string
	Offensive string
}

type preferences struct {
	InvertColorScheme bool
}

func (s state) Statement() *statement.Statement {
	return s.statement
}

func (s state) Categories() categories {
	c := categories{}

	if s.categories.Has(category.Harmless) {
		c.Harmless, _ = translate.TranslateCategory(category.Harmless, s.language)
	}

	if s.categories.Has(category.Delicate) {
		c.Delicate, _ = translate.TranslateCategory(category.Delicate, s.language)
	}

	if s.categories.Has(category.Offensive) {
		c.Offensive, _ = translate.TranslateCategory(category.Offensive, s.language)
	}

	return c
}

func (s state) Language() string {
	return s.language.String()
}

func (s state) Preferences() preferences {
	return *s.preferences
}

func (s state) ContentLocation() string {
	return "/statements/" + s.statement.ID.String() + "?language=" + s.language.String()
}
