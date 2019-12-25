package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/neverhaveiever-io/api/internal/app/router"
	"github.com/neverhaveiever-io/api/internal/database"
	"github.com/neverhaveiever-io/api/internal/statement"
	"github.com/spf13/viper"
)

func main() {
	defer database.C.Close()

	user := viper.GetString("admin_user")
	pass := viper.GetString("admin_pass")

	if user == "" {
		panic("no admin user")
	}

	if pass == "" {
		panic("no admin password")
	}

	router.Init(gin.BasicAuth(gin.Accounts{
		user: pass,
	}))
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

	// create category type
	database.C.Exec("create type category as enum ('harmless', 'delicate', 'offensive')")

	// migrate model
	database.C.AutoMigrate(&statement.Statement{})
}
