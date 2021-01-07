package main

import (
	"github.com/nhie-io/api/internal/app/router"
	"github.com/nhie-io/api/internal/cache"
	"github.com/nhie-io/api/internal/database"
	"github.com/nhie-io/api/internal/history"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"github.com/spf13/viper"
)

func main() {
	router.Init()
}

func init() {
	viper.SetEnvPrefix("NHIE")
	viper.AutomaticEnv()

	// initialize db connection
	if err := database.Init(); err != nil {
		panic("unable to connect to database: " + err.Error())
	}

	// initialize translate
	if err := translate.Init(); err != nil {
		panic("failed to initialize translations: " + err.Error())
	}

	// initialize cache
	cache.Init()

	// initialize history
	history.Init()

	// create category type
	database.C.Exec("create type category as enum ('harmless', 'delicate', 'offensive')")

	// migrate model
	if err := database.C.AutoMigrate(&statement.Statement{}); err != nil {
		panic("failed to auto migrate models: " + err.Error())
	}
}
