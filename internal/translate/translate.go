package translate

import (
	translate "cloud.google.com/go/translate/apiv3"
	"context"
	"fmt"
	"github.com/google/uuid"
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
var m = language.NewMatcher([]language.Tag{
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
		// TODO: handle gracefully
		panic("gc project not set")
	}

	C = &TranslationClient{
		c:          c,
		ctx:        ctx,
		sourceLang: language.AmericanEnglish,
		project:    viper.GetString("translate_project"),
		location:   viper.GetString("translate_location"),
		model:      viper.GetString("translate_model"),
	}

	return nil
}

func MatchTags(inputs ...string) []language.Tag {
	var tags []language.Tag

	for _, input := range inputs {
		tag, _, confidence := m.Match(language.Make(input))

		// TODO: error if no match
		if confidence == language.Exact {
			tags = append(tags, tag)
		}
	}

	return tags
}

func (tc *TranslationClient) Translate(uuid uuid.UUID, s string, tag language.Tag) (string, error) {

	ttr, err := retrieveFromCache(uuid, tag, tc.model)

	if err == nil && len(ttr.Translations) > 0 {
		return ttr.Translations[0].TranslatedText, nil
	}

	req := &translatepb.TranslateTextRequest{
		Contents:           []string{s},
		MimeType:           "text/plain",
		SourceLanguageCode: tc.sourceLang.String(),
		TargetLanguageCode: tag.String(),
		Parent:             fmt.Sprintf("projects/%s/locations/%s", tc.project, tc.location),
		Model:              fmt.Sprintf("projects/%s/locations/%s/models/%s", tc.project, tc.location, tc.model),
	}

	resp, err := tc.c.TranslateText(tc.ctx, req)

	if err != nil {
		panic("failed to translate text: " + err.Error())
	}

	if len(resp.Translations) == 0 {
		// TODO: handle gracefully
		panic("no translation received")
	}

	// TODO: handle error
	_ = storeInCache(uuid, tag, tc.model, resp)
	return resp.Translations[0].TranslatedText, nil
}

func BulkTranslate() {
	// TODO: implement
}
