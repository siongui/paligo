package main

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/siongui/gopalilib/tpkutil"
)

func main() {
	websiteDir := flag.String("websiteDir", "", "output dir of website")
	supportedLocales := flag.String("supportedLocales", "", "supported locales on website, separated by ,")
	flag.Parse()

	fmt.Println(*websiteDir)
	err := tpkutil.SymlinkToRootIndexHtml(*websiteDir, "romn")
	if err != nil {
		panic(err)
	}
	locales := strings.Split(*supportedLocales, ",")
	for _, locale := range locales {
		dir := path.Join(*websiteDir, locale)
		fmt.Println(dir)
		err := tpkutil.SymlinkToRootIndexHtml(dir, "romn")
		if err != nil {
			panic(err)
		}
	}
}
