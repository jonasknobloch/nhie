package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var C *gorm.DB

func Init(dsn string) error {
	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		return err
	}

	C = client

	return nil
}
