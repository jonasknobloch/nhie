package translate

import (
	translate "cloud.google.com/go/translate/apiv3"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/cache"
	"github.com/spf13/viper"
	"golang.org/x/text/language"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

type TranslationClient struct {
	c          *translate.TranslationClient
	ctx        context.Context
	sourceLang language.Tag
	project    string
	location   string
	model      string
}

var C *TranslationClient
var SourceLanguage = language.AmericanEnglish
var m = language.NewMatcher([]language.Tag{
	SourceLanguage,
	language.German,
	language.Spanish,
})

func Init() error {
	ctx := context.Background()
	c, err := translate.NewTranslationClient(ctx)

	if err != nil {
		return err
	}

	viper.SetDefault("translate_location", "global")
	viper.SetDefault("translate_model", "general/nmt")

	if !viper.IsSet("translate_project") {
		return errors.New("gc project not set")
	}

	C = &TranslationClient{
		c:          c,
		ctx:        ctx,
		sourceLang: SourceLanguage,
		project:    viper.GetString("translate_project"),
		location:   viper.GetString("translate_location"),
		model:      viper.GetString("translate_model"),
	}

	return nil
}

func MatchTags(inputs ...string) ([]language.Tag, error) {
	var tags []language.Tag

	for _, input := range inputs {
		tag, _, confidence := m.Match(language.Make(input))

		if confidence == language.Exact {
			tags = append(tags, tag)
			continue
		}

		// parse returns ValueError if well formed
		parsedTag, err := language.Parse(input)

		// well formed but unknown language tag
		if _, ok := err.(language.ValueError); ok {
			return nil, newError(newMatchingError(ErrUnknownLanguageTag, input, parsedTag))
		}

		// invalid language tag
		if err != nil {
			return nil, newError(newMatchingError(ErrInvalidLanguageTag, input, language.Tag{}))
		}

		return nil, newError(newMatchingError(ErrUnsupportedLanguage, input, parsedTag))
	}

	return tags, nil
}

func (tc *TranslationClient) Translate(uuid uuid.UUID, s string, tag language.Tag) (string, error) {
	var ttr *translatepb.TranslateTextResponse
	var cacheErr error
	var fetchErr error

	var t string

	// translation cached
	if t, cacheErr = retrieveFromCache(uuid, tag, tc.model); cacheErr == nil {
		return t, nil
	}

	ttr, fetchErr = tc.fetchFromApi(s, tag)

	// failed fetch
	if fetchErr != nil {
		return "", newError(fetchErr)
	}

	// verify a translation is present
	if len(ttr.Translations) == 0 {
		return "", newError(ErrNoTranslationReceived)
	}

	t = ttr.Translations[0].TranslatedText

	// store if possible
	if errors.Is(cacheErr, cache.ErrKeyNotFound) {
		cacheErr = storeInCache(uuid, tag, tc.model, t)
	}

	// unwrapped cache error might be nil
	if errors.Unwrap(cacheErr) != nil {
		return ttr.Translations[0].TranslatedText, newError(cacheErr)
	}

	return t, nil
}

func (tc *TranslationClient) fetchFromApi(s string, tag language.Tag) (*translatepb.TranslateTextResponse, error) {
	req := &translatepb.TranslateTextRequest{
		Contents:           []string{s},
		MimeType:           "text/plain",
		SourceLanguageCode: tc.sourceLang.String(),
		TargetLanguageCode: tag.String(),
		Parent:             fmt.Sprintf("projects/%s/locations/%s", tc.project, tc.location),
		Model:              fmt.Sprintf("projects/%s/locations/%s/models/%s", tc.project, tc.location, tc.model),
	}

	return tc.c.TranslateText(tc.ctx, req)
}

func BulkTranslate() {
	// TODO: implement
}
