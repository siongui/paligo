package main

import (
	"flag"
	"os"

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
	isdev := flag.Bool("isdev", false, "Is development?")
	flag.Parse()

	if *action == "symlink" {
		err := dicutil.SymlinkToRootIndexHtml(wordJsonDir, websiteDir)
		if err != nil {
			panic(err)
		}
	}

	if *action == "html" {
		data := dicutil.TemplateData{
			SiteUrl:     "https://dictionary.online-dhamma.net",
			TipitakaURL: "http://tipitaka.online-dhamma.net",
			OgImage:     "https://upload.wikimedia.org/wikipedia/commons/d/df/Dharma_Wheel.svg",
			OgUrl:       "https://dictionary.online-dhamma.net/",
			OgLocale:    "en_US",
		}

		if *isdev {
			data.SiteUrl = ""
		}

		err := dicutil.CreateHTML(os.Stdout, "index.html", &data, localeDir, htmlTemplateDir)
		if err != nil {
			panic(err)
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
