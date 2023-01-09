package translate

import (
	"context"
	"errors"
	"github.com/bounoable/deepl"
	"github.com/jonasknobloch/nhie/internal/database"
	"github.com/jonasknobloch/nhie/internal/statement"
	"golang.org/x/text/language"
	"time"
)

var (
	ErrInvalidLanguageTag  = errors.New("invalid language tag")
	ErrUnsupportedLanguage = errors.New("unsupported language")
)

var deeplLanguage = map[language.Tag]deepl.Language{
	language.English: deepl.EnglishAmerican,
	language.German:  deepl.German,
	language.Swedish: deepl.Swedish,
}

var client *deepl.Client

func Init(authKey, baseURL string) {
	client = deepl.New(authKey, deepl.BaseURL(baseURL))
}

func MatchTag(input string) (language.Tag, error) {
	tag, err := language.Parse(input)

	if err != nil {
		return language.Und, ErrInvalidLanguageTag
	}

	if _, ok := deeplLanguage[tag]; !ok {
		return language.Und, ErrUnsupportedLanguage
	}

	return tag, nil
}

func TranslateStatement(s *statement.Statement, t language.Tag) error {
	l, ok := deeplLanguage[t]

	if !ok {
		return ErrUnsupportedLanguage
	}

	translation, _, err := client.Translate(context.TODO(), s.Statement, l)

	if err != nil {
		return err
	}

	now := time.Now()

	if err := database.C.Exec(`INSERT INTO translations (statement_id, language, translation, created_at, updated_at) VALUES (?, ?, ?, ?, ?);`, s.ID, t.String(), translation, now, now).Error; err != nil {
		return err
	}

	return nil
}

func TranslateMissing(t language.Tag) error {
	statements := make([]statement.Statement, 0)

	_, ok := deeplLanguage[t]

	if !ok {
		return ErrUnsupportedLanguage
	}

	if err := database.C.Raw(`SELECT id, statement, category FROM statements WHERE NOT EXISTS (SELECT statement_id FROM translations WHERE language = ? AND statements.id = translations.statement_id);`, t.String()).Scan(&statements).Error; err != nil {
		return err
	}

	for _, s := range statements {
		if err := TranslateStatement(&s, t); err != nil {
			return err
		}
	}

	return nil
}
