package translate

import (
	"github.com/nhie-io/api/internal/category"
	"golang.org/x/text/language"
)

var categories = map[category.Category]map[language.Tag]string{
	category.Harmless: {
		language.English: "harmless",
		language.German:  "harmlos",
	},
	category.Delicate: {
		language.English: "delicate",
		language.German:  "delikat",
	},
	category.Offensive: {
		language.English: "offensive",
		language.German:  "offensiv",
	},
}

func TranslateCategory(category category.Category, tag language.Tag) (string, bool) {
	t, ok := categories[category][tag]

	if !ok {
		return string(category), false
	}

	return t, true
}
