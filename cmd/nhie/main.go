package main

import (
	"errors"
	"fmt"
	"github.com/nhie-io/api/internal/application"
	"github.com/nhie-io/api/internal/database"
	"log"
	"os"
)

const PostgresDSNEnv = "NHIE_POSTGRES_DSN"

func main() {
	if db, err := database.C.DB(); err == nil {
		defer db.Close()
	}

	if err := application.Init(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	dsn, ok := os.LookupEnv(PostgresDSNEnv)

	if !ok {
		log.Fatal(envNotSetError(PostgresDSNEnv))
	}

	if err := database.Init(dsn); err != nil {
		log.Fatal(err)
	}
}

func envNotSetError(env string) error {
	return errors.New(fmt.Sprintf("required environment variabale \"%s\" not set", env))
}
