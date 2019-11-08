package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var C *gorm.DB

func Init(connString string) {
	var err error
	C, err = gorm.Open("postgres", connString)

	if err != nil {
		panic("unable to connect to database: " + err.Error())
	}
}
