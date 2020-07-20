package main

import (
	"flag"
	"fmt"
	"path"
	"strings"

	"github.com/siongui/gopalilib/dicutil"
	"github.com/siongui/gopherjs-i18n/tool"
)

const websiteDir = "website"
const localeDir = "locale"
const poDomain = "messages"
const htmlTemplateDir = "theme/template"

const bookCSV = "data/dictionary/dict-books.csv"
const wordCSV1 = "data/dictionary/dict_words_1.csv"
const wordCSV2 = "data/dictionary/dict_words_2.csv"

const outBookJSON = websiteDir + "/bookIdAndInfos.json"
const wordJsonDir = websiteDir + "/json/"

const trieDataPath = websiteDir + "/strie.txt"
const trieNodeCountPath = websiteDir + "/strie_node_count.txt"
const rankDirectoryDataPath = websiteDir + "/rd.txt"

const poJsonPath = websiteDir + "/po.json"

func main() {
	action := flag.String("action", "", "What kind of action?")
	websiteDir := flag.String("websiteDir", "", "output dir of website")
	supportedLocales := flag.String("supportedLocales", "", "supported locales on website, separated by ,")
	flag.Parse()

	if *action == "symlink" {
		fmt.Println(*websiteDir)
		err := dicutil.SymlinkToRootIndexHtml(*websiteDir)
		if err != nil {
			panic(err)
		}
		locales := strings.Split(*supportedLocales, ",")
		for _, locale := range locales {
			dir := path.Join(*websiteDir, locale)
			fmt.Println(dir)
			err := dicutil.SymlinkToRootIndexHtml(dir)
			if err != nil {
				panic(err)
			}
		}
	}

	if *action == "parsebooks" {
		dicutil.ParseDictionayBookCSV(bookCSV, outBookJSON)
	}

	if *action == "parsewords" {
		dicutil.ParseDictionayWordCSV(wordCSV1, wordCSV2, wordJsonDir)
	}

	if *action == "triebuild" {
		dicutil.BuildSuccinctTrie(wordJsonDir, trieDataPath, trieNodeCountPath, rankDirectoryDataPath)
	}

	if *action == "po2json" {
		po2json.PO2JSON(poDomain, localeDir, poJsonPath)
	}
}
