package translate

import (
	"golang.org/x/text/language"
)

func MatchTag(input string) (language.Tag, error) {
	parsedTag, err := language.Parse(input)

	// well formed but unknown language tag
	if _, ok := err.(language.ValueError); ok {
		return language.Und, newError(newMatchingError(ErrUnknownLanguageTag, input, parsedTag))
	}

	// invalid language tag
	if err != nil {
		return language.Und, newError(newMatchingError(ErrInvalidLanguageTag, input, language.Tag{}))
	}

	tag, _, confidence := m.Match(parsedTag)

	if confidence == language.Exact {
		return tag, nil
	}

	return language.Und, newError(newMatchingError(ErrUnsupportedLanguage, input, parsedTag))
}
