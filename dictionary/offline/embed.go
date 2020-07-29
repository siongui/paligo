// +build ignore

package main

import (
	"flag"

	"github.com/siongui/goef"
)

func main() {
	websiteDir := flag.String("websiteDir", "", "output dir of website")
	flag.Parse()

	err := goef.GenerateGoPackage("main", *websiteDir, "offline/data.go")
	if err != nil {
		panic(err)
	}
}
