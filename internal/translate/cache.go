package translate

import (
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/cache"
	"golang.org/x/text/language"
)

func retrieveFromCache(uuid uuid.UUID, tag language.Tag, model string) (string, error) {
	// translate:UUID:tag:model
	key := cache.Key{"translate", uuid.String(), tag.String(), model}
	t, err := cache.Retrieve(key)

	if err != nil {
		return "", err
	}

	return t, nil
}

func storeInCache(uuid uuid.UUID, tag language.Tag, model string, t string) error {
	// translate:UUID:tag:model
	key := cache.Key{"translate", uuid.String(), tag.String(), model}
	return cache.Store(key, t, 0)
}

func ClearCache(uuid uuid.UUID) {
	cache.Clear(cache.Key{"translate", uuid.String(), "*"})
}
