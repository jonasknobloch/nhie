package history

import (
	"github.com/google/uuid"
	"github.com/neverhaveiever-io/api/internal/cache"
	"github.com/neverhaveiever-io/api/internal/statement"
	"time"
)

func Exists(gameID uuid.UUID, statement *statement.Statement) (bool, error) {
	k := cache.Key{"history", gameID.String(), statement.ID.String()}
	e, err := cache.Exists(k)

	if err != nil {
		return false, err
	}

	return e, nil
}

func Add(gameID uuid.UUID, statement *statement.Statement) error {
	k := cache.Key{"history", gameID.String(), statement.ID.String()}
	return cache.Store(k, time.Now().String(), 4*time.Hour) // TODO: use config value
}
