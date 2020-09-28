package history

import (
	"github.com/google/uuid"
	"github.com/nhie-io/api/internal/cache"
	"github.com/nhie-io/api/internal/statement"
	"github.com/spf13/viper"
	"time"
)

var MaxTries int
var ttl time.Duration

func Init() {
	viper.SetDefault("history_max_tries", 5)
	viper.SetDefault("history_ttl", int64(4*time.Hour))

	MaxTries = viper.GetInt("history_max_tries")
	ttl = time.Duration(viper.GetInt64("history_ttl"))
}

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
	return cache.Store(k, time.Now().String(), ttl)
}
