package web

import (
	"embed"
	"io/fs"
)

//go:embed build
var build embed.FS
var Build fs.FS

//go:embed index.html
var Index string

func init() {
	sub, err := fs.Sub(build, "build")

	if err != nil {
		panic(err)
	}

	Build = sub
}
