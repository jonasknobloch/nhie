package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var C *gorm.DB

func Init(dsn string) error {
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Error),
		PrepareStmt: true,
	})

	if err != nil {
		return err
	}

	C = client

	return nil
}
