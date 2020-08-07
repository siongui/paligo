package main

// Embed Tipiṭaka ToC JSON in Go code

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/tpkutil"
	"github.com/siongui/gopalilib/util"
)

func SaveTipitakaTreeJson(xmldir, outputdir string) {
	fmt.Println(xmldir)
	t, err := tpkutil.BuildTipitakaTree(xmldir)
	if err != nil {
		panic(err)
	}
	//fmt.Println(t)
	util.SaveJsonFile(t, path.Join(outputdir, "tpktoc.json"))
}

func main() {
	xmlDir := flag.String("xmlDir", "", "Tipiṭaka XML dir")
	flag.Parse()

	dir, err := ioutil.TempDir("", "tocjson")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir) // clean up

	SaveTipitakaTreeJson(*xmlDir, dir)

	err = goef.GenerateGoPackagePlainText("main", dir, "gopherjs/data.go")
	if err != nil {
		panic(err)
	}
}
