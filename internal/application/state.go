package application

import (
	"github.com/nhie-io/api/internal/category"
	"github.com/nhie-io/api/internal/statement"
	"golang.org/x/text/language"
)

type state struct {
	statement  *statement.Statement
	categories *category.Selection
	language   language.Tag
}

type categories struct {
	Harmless  bool
	Delicate  bool
	Offensive bool
}

func (s state) Statement() *statement.Statement {
	return s.statement
}

func (s state) Categories() categories {
	return categories{
		Harmless:  s.categories.Has(category.Harmless),
		Delicate:  s.categories.Has(category.Delicate),
		Offensive: s.categories.Has(category.Offensive),
	}
}

func (s state) Language() string {
	return s.language.String()
}

func (s state) ContentLocation() string {
	return "/statements/" + s.statement.ID.String()
}
