package cache

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var C *redis.Client

func Init() {
	viper.SetDefault("redis_host", "localhost")
	viper.SetDefault("redis_port", 6379)
	viper.SetDefault("redis_pass", "")

	C = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis_host") + ":" + viper.GetString("redis_port"),
		Password: viper.GetString("redis_password"),
		DB:       0,
	})
}
