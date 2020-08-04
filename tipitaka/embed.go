package main

// Embed Tipiṭaka ToC JSON in Go code

import (
	"flag"
	"fmt"
	"io/ioutil"
	"path"
	"strconv"

	"github.com/siongui/go-succinct-data-structure-trie"
	"github.com/siongui/goef"
	"github.com/siongui/gopalilib/dicutil"
	"github.com/siongui/gopalilib/tpkutil"
	"github.com/siongui/gopalilib/util"
	"github.com/siongui/gopaliwordvfs"
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

func SaveBookIdAndInfosJson(datarepodir, outputdir string) {
	bookCsv := path.Join(datarepodir, "dictionary/dict-books.csv")
	fmt.Println(bookCsv)
	dicutil.ParseDictionayBookCSV(bookCsv, path.Join(outputdir, "BookIdAndInfos.json"))
}

func SaveTrieJson(words []string, outputdir string) {
	// set alphabet of words
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	// encode: build succinct trie
	te := bits.Trie{}
	te.Init()

	for i, word := range words {
		fmt.Println(i, word)
		te.Insert(word)
	}

	// encode: trie encoding
	teData := te.Encode()
	//fmt.Println(teData)
	err := ioutil.WriteFile(path.Join(outputdir, "trie-data.txt"), []byte(teData), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(te.GetNodeCount())
	err = ioutil.WriteFile(path.Join(outputdir, "trie-node-count.txt"), []byte(strconv.Itoa(int(te.GetNodeCount()))), 0644)
	if err != nil {
		panic(err)
	}

	// encode: build cache for quick lookup
	rd := bits.CreateRankDirectory(teData, te.GetNodeCount()*2+1, bits.L1, bits.L2)
	//fmt.Println(rd.GetData())
	err = ioutil.WriteFile(path.Join(outputdir, "trie-rank-directory-data.txt"), []byte(rd.GetData()), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	xmlDir := flag.String("xmlDir", "", "Tipiṭaka XML dir")
	dataDir := flag.String("dataDir", "", "website data dir")
	dataRepoDir := flag.String("dataRepoDir", "", "dir of data repo which contains pali data")
	flag.Parse()

	SaveTipitakaTreeJson(*xmlDir, *dataDir)
	SaveBookIdAndInfosJson(*dataRepoDir, *dataDir)
	SaveTrieJson(gopaliwordvfs.MapKeys(), *dataDir)

	err := goef.GenerateGoPackagePlainText("main", *dataDir, "gopherjs/data.go")
	if err != nil {
		panic(err)
	}
}
