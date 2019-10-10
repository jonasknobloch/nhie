package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/neverhaveiever-io/api/database"
	"github.com/neverhaveiever-io/api/models"
	"github.com/neverhaveiever-io/api/routers"
	"github.com/spf13/viper"
)

func main() {
	defer database.Connection.Close()

	user := viper.GetString("admin_user")
	pass := viper.GetString("admin_pass")

	if user == "" {
		panic("no admin user")
	}

	if pass == "" {
		panic("no admin password")
	}

	router := routers.InitRouter(gin.BasicAuth(gin.Accounts{
		user: pass,
	}))

	err := router.Run()

	if err != nil {
		panic("unable initialize router")
	}
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
		viper.GetString("db_user"),
	)

	var err error
	database.Connection, err = gorm.Open("postgres", connString)

	if err != nil {
		panic("unable to connect to database: " + err.Error())
	}

	// migrate model
	database.Connection.AutoMigrate(&models.Statement{})
}
