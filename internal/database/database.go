package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var C *gorm.DB

func Init() error {
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 5432)
	viper.SetDefault("db_name", "neverhaveiever")
	viper.SetDefault("db_user", "neverhaveiever")
	viper.SetDefault("db_pass", nil)

	connection := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		viper.GetString("db_host"),
		viper.GetString("db_port"),
		viper.GetString("db_name"),
		viper.GetString("db_user"),
		viper.GetString("db_pass"),
	)

	if client, err := gorm.Open(postgres.Open(connection), &gorm.Config{}); err != nil {
		return err
	} else {
		C = client
	}

	return nil
}
