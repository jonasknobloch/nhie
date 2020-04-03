package translate

import (
	"github.com/google/uuid"
	"github.com/neverhaveiever-io/api/internal/cache"
	"golang.org/x/text/language"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

func retrieveFromCache(uuid uuid.UUID, tag language.Tag, model string) (*translatepb.TranslateTextResponse, error) {
	var ttr translatepb.TranslateTextResponse

	// translate:UUID:tag:model
	key := cache.Key{"translate", uuid.String(), tag.String(), model}
	err := cache.Retrieve(key, &ttr)

	if err != nil {
		return nil, err
	}

	return &ttr, nil
}

func storeInCache(uuid uuid.UUID, tag language.Tag, model string, ttr *translatepb.TranslateTextResponse) error {
	// translate:UUID:tag:model
	key := cache.Key{"translate", uuid.String(), tag.String(), model}
	return cache.Store(key, ttr, 0)
}

func ClearCache(uuid uuid.UUID) {
	cache.Clear(cache.Key{"translate", uuid.String(), "*"})
}
