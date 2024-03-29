package application

import (
	"github.com/google/uuid"
	"github.com/jonasknobloch/nhie/internal/category"
	"github.com/jonasknobloch/nhie/internal/translate"
	"golang.org/x/text/language"
	"net/http"
	"strconv"
)

func queryLanguage(request *http.Request) (language.Tag, bool) {
	val := request.URL.Query().Get("language")

	if val == "" {
		return translate.SourceLanguage, true
	}

	tag, err := translate.MatchTag(val)

	if err != nil {
		return language.Tag{}, false
	}

	return tag, true
}

func queryCategories(request *http.Request) (*category.Selection, bool) {
	query := request.URL.Query()

	selection := category.NewSelection()

	for _, val := range query["category"] {
		if c, ok := category.Match(val); !ok {
			return selection, false
		} else {
			selection.Add(c)
		}
	}

	if selection.Empty() {
		selection.Add(category.Harmless)
		selection.Add(category.Delicate)
		selection.Add(category.Offensive)
	}

	return selection, true
}

func queryStatementID(request *http.Request) (uuid.UUID, bool) {
	val := request.URL.Query().Get("statement_id")

	if val == "" {
		return uuid.UUID{}, false
	}

	id, err := uuid.Parse(val)

	if err != nil {
		return uuid.UUID{}, false
	}

	return id, true
}

func queryInvertColorScheme(request *http.Request) (bool, bool) {
	val := request.URL.Query().Get("invert_color_scheme")

	if val == "" {
		return false, false
	}

	b, err := strconv.ParseBool(val)

	if err != nil {
		return false, false
	}

	return b, true
}
