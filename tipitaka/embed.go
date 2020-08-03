package main

// Embed Tipiṭaka ToC JSON in Go code

import (
	"flag"
	"fmt"
	"path"

	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/dicutil"
	"github.com/siongui/gopalilib/tpkutil"
	"github.com/siongui/gopalilib/util"
)

func main() {
	xmlDir := flag.String("xmlDir", "", "Tipiṭaka XML dir")
	dataDir := flag.String("dataDir", "", "website data dir")
	dataRepoDir := flag.String("dataRepoDir", "", "dir of data repo which contains pali data")
	flag.Parse()

	fmt.Println(*xmlDir)
	t, err := tpkutil.BuildTipitakaTree(*xmlDir)
	if err != nil {
		panic(err)
	}
	//fmt.Println(t)

	util.SaveJsonFile(t, path.Join(*dataDir, "tpktoc.json"))

	bookCsv := path.Join(*dataRepoDir, "dictionary/dict-books.csv")
	fmt.Println(bookCsv)
	dicutil.ParseDictionayBookCSV(bookCsv, path.Join(*dataDir, "BookIdAndInfos.json"))

	err = goef.GenerateGoPackagePlainText("main", *dataDir, "gopherjs/data.go")
	if err != nil {
		panic(err)
	}
}
