package main

import (
	"fmt"
	"github.com/jonasknobloch/nhie/internal/database"
	"github.com/jonasknobloch/nhie/internal/translate"
	"golang.org/x/text/language"
	"log"
	"os"
)

const (
	PostgresDSNEnv  = "NHIE_POSTGRES_DSN"
	DeeplAuthKeyEnv = "NHIE_DEEPL_AUTH_KEY"
	DeeplBaseURLEnv = "NHIE_DEEPL_BASE_URL"
)

func main() {
	if db, err := database.C.DB(); err == nil {
		defer db.Close()
	}

	err := translate.TranslateMissing(language.German)

	if err != nil {
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

	authKey, ok := os.LookupEnv(DeeplAuthKeyEnv)

	if !ok {
		log.Fatal(envNotSetError(DeeplAuthKeyEnv))
	}

	baseURL, ok := os.LookupEnv(DeeplBaseURLEnv)

	if !ok {
		log.Fatal(envNotSetError(DeeplBaseURLEnv))
	}

	translate.Init(authKey, baseURL)
}

func envNotSetError(env string) error {
	return fmt.Errorf("required environment variabale \"%s\" not set", env)
}
