package main

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/siongui/gopalilib/dicutil"
	"github.com/siongui/gopaliwordvfs"
)

func main() {
	action := flag.String("action", "", "What kind of action?")
	websiteDir := flag.String("websiteDir", "", "output dir of website")
	supportedLocales := flag.String("supportedLocales", "", "supported locales on website, separated by ,")
	flag.Parse()

	if *action == "symlink" {
		fmt.Println(*websiteDir)
		err := dicutil.SymlinkToRootIndexHtml(*websiteDir, gopaliwordvfs.MapKeys())
		if err != nil {
			panic(err)
		}
		locales := strings.Split(*supportedLocales, ",")
		for _, locale := range locales {
			dir := path.Join(*websiteDir, locale)
			fmt.Println(dir)
			err := dicutil.SymlinkToRootIndexHtml(dir, gopaliwordvfs.MapKeys())
			if err != nil {
				panic(err)
			}
		}
	}
}
