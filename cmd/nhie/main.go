package main

import (
	"fmt"
	"github.com/neverhaveiever-io/api/internal/app/router"
	"github.com/neverhaveiever-io/api/internal/cache"
	"github.com/neverhaveiever-io/api/internal/database"
	"github.com/neverhaveiever-io/api/internal/statement"
	"github.com/neverhaveiever-io/api/internal/translate"
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

	// TODO: handle error
	// initialize translate
	_ = translate.Init()

	// initialize cache
	cache.Init()

	// create category type
	database.C.Exec("create type category as enum ('harmless', 'delicate', 'offensive')")

	// migrate model
	database.C.AutoMigrate(&statement.Statement{})
}
