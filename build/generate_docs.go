// +build ignore

package main

import (
	"go/build"
	"log"
	"os"
	"path"

	"github.com/spf13/cobra/doc"
	"github.com/willjcj/param/cmd"
)

func main() {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}

	docsPath := path.Join(gopath, "src", "github.com", "willjcj", "param", "docs")

	err := doc.GenMarkdownTree(cmd.RootCmd, docsPath)
	if err != nil {
		log.Fatal(err)
	}
}
