package translate

import (
	"golang.org/x/text/language"
)

var Languages = []language.Tag{
	language.English,
	language.German,
}

var SourceLanguage = Languages[0]

var matcher = language.NewMatcher(Languages)

func MatchLanguage(preferences []language.Tag) language.Tag {
	_, index, confidence := matcher.Match(preferences...)

	if confidence == language.No {
		return Languages[0]
	}

	return Languages[index]
}
