package main

import (
	"fmt"
	"github.com/evanw/esbuild/pkg/api"
	"io"
	"log"
)
import "os"

func main() {
	onRebuild := func(result api.BuildResult) {
		if len(result.Errors) > 0 {
			fmt.Printf("Build failed: %d errors\n", len(result.Errors))
		} else {
			fmt.Printf("Build succeeded: %d warnings\n", len(result.Warnings))
		}
	}

	result := api.Build(api.BuildOptions{
		EntryPoints:       []string{"web/app.js"},
		Outdir:            "web/build",
		MinifyWhitespace:  true,
		MinifyIdentifiers: false,
		Write:             true,
		Bundle:            true,
		Watch: &api.WatchMode{
			OnRebuild: onRebuild,
		},
	})

	if len(result.Errors) > 0 {
		for _, m := range result.Errors {
			fmt.Println(m)
		}

		os.Exit(1)
	}

	if err := copyFile("web/static/mojito.svg", "web/build/mojito.svg"); err != nil {
		log.Fatal(err)
	}

	if err := copyFile("web/static/beer.svg", "web/build/beer.svg"); err != nil {
		log.Fatal(err)
	}

	if err := copyFile("web/static/wine.svg", "web/build/wine.svg"); err != nil {
		log.Fatal(err)
	}

	if err := copyFile("web/static/favicon-16x16.png", "web/build/favicon-16x16.png"); err != nil {
		log.Fatal(err)
	}

	if err := copyFile("web/static/favicon-32x32.png", "web/build/favicon-32x32.png"); err != nil {
		log.Fatal(err)
	}

	if err := copyFile("web/static/favicon-96x96.png", "web/build/favicon-96x96.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Watching...")

	<-make(chan bool)
}

func copyFile(src, dst string) error {
	s, err := os.Open(src)

	if err != nil {
		return err
	}

	defer s.Close()

	c, err := os.Create(dst)

	if err != nil {
		return err
	}

	defer c.Close()

	_, err = io.Copy(c, s)

	return err
}
