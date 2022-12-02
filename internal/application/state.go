package application

import (
	"github.com/jonasknobloch/nhie/internal/category"
	"github.com/jonasknobloch/nhie/internal/statement"
	"github.com/jonasknobloch/nhie/internal/translate"
	"golang.org/x/text/language"
)

type state struct {
	statement   *statement.Statement
	categories  *category.Selection
	language    language.Tag
	preferences *preferences
	config      *config
}

type categories struct {
	Harmless  string
	Delicate  string
	Offensive string
}

type preferences struct {
	InvertColorScheme bool
}

type config struct {
	WebHost   string
	ApiHost   string
	URLScheme string
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

func (s state) Config() config {
	return *s.config
}

func (s state) ContentLocation() string {
	return "/statements/" + s.statement.ID.String() + "?language=" + s.language.String()
}
