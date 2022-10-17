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
	fmt.Println("      _    _     ")
	fmt.Println(" _ _ | |_ (_)___ ")
	fmt.Println("| ' \\| ' \\| / -_)")
	fmt.Println("|_||_|_||_|_\\___|")
	fmt.Println("31 38 31 32 31 35")

	fmt.Print("\n")

	if db, err := database.C.DB(); err != nil {
		log.Fatal(err)
	} else {
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
