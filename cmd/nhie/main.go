package main

import (
	"fmt"
	"github.com/nhie-io/api/internal/app/router"
	"github.com/nhie-io/api/internal/cache"
	"github.com/nhie-io/api/internal/database"
	"github.com/nhie-io/api/internal/history"
	"github.com/nhie-io/api/internal/statement"
	"github.com/nhie-io/api/internal/translate"
	"github.com/spf13/viper"
)

func main() {
	defer database.C.Close()

	router.Init()
}

func init() {
	viper.SetEnvPrefix("NHIE")
	viper.AutomaticEnv()

	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 5432) // default port postgres port
	viper.SetDefault("db_name", "neverhaveiever")
	viper.SetDefault("db_user", "neverhaveiever")
	viper.SetDefault("db_pass", nil)

	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_name"),
		viper.GetString("db_user"),
		viper.GetString("db_pass"),
	)

	// initialize db connection
	database.Init(connString)

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
	database.C.AutoMigrate(&statement.Statement{})
}
