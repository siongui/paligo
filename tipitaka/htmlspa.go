// Create the only index.html for Single Page Application (SPA), and 404.html to
// be served if not found.
package main

import (
	"flag"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/siongui/gopalilib/util"
	"github.com/siongui/gtmpl"
)

// Template data for webpage of Pāli Dictionary
type TemplateData struct {
	SiteUrl       string
	DictionaryURL string
	OgImage       string
	OgUrl         string
	OgLocale      string
}

func CreateIndexAnd404(tmpl *template.Template, data TemplateData, locale, websiteDir string) (err error) {
	pi := path.Join(websiteDir, locale, "index.html")
	p4 := path.Join(websiteDir, locale, "404.html")

	/*
		println(pi)
		println(p4)
		println(data.OgLocale)
	*/

	util.CreateDirIfNotExist(pi)
	findex, err := os.Create(pi)
	if err != nil {
		return
	}
	defer findex.Close()

	util.CreateDirIfNotExist(p4)
	f404, err := os.Create(p4)
	if err != nil {
		return
	}
	defer f404.Close()

	err = tmpl.ExecuteTemplate(findex, "index.html", &data)
	if err != nil {
		return
	}
	return tmpl.ExecuteTemplate(f404, "404.html", &data)
}

func main() {
	siteconffile := flag.String("siteconf", "", "JSON config file for website")
	websiteDir := flag.String("websiteDir", "", "output dir of website")
	htmlTemplateDir := flag.String("htmlTemplateDir", "", "html template dir")
	localeDir := flag.String("localeDir", "", "locale translation dir")
	flag.Parse()

	siteconf, err := util.LoadJsonConfig(*siteconffile)
	if err != nil {
		panic(err)
	}

	supportedLocales := siteconf["supportedLocales"]
	data := TemplateData{
		SiteUrl:       siteconf["SiteUrl"],
		DictionaryURL: siteconf["DictionaryURL"],
		OgImage:       siteconf["OgImage"],
		OgUrl:         siteconf["OgUrl"],
		OgLocale:      "",
	}

	tmpl, err := gtmpl.ParseDirectoryTree("messages", *localeDir, *htmlTemplateDir)
	if err != nil {
		panic(err)
	}

	err = CreateIndexAnd404(tmpl, data, "", *websiteDir)

	sl := strings.Split(supportedLocales, ",")
	for _, locale := range sl {
		gtmpl.SetLanguage(locale)
		data.OgLocale = locale
		err = CreateIndexAnd404(tmpl, data, locale, *websiteDir)
		if err != nil {
			panic(err)
		}
	}
}
