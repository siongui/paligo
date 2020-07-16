// Create the only index.html for Single Page Application (SPA), and 404.html to
// be served if not found.
package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"os"
	"path"

	"github.com/siongui/gopalilib/util"
	"github.com/siongui/gtmpl"
)

// Template data for webpage of Pāli Dictionary
type TemplateData struct {
	SiteUrl     string
	TipitakaURL string
	OgImage     string
	OgUrl       string
	OgLocale    string
}

func LoadJsonConfig(fp string) (conf map[string]string, err error) {
	f, err := os.Open(fp)
	if err != nil {
		return
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	err = dec.Decode(&conf)
	return
}

func CreateIndexAnd404(tmpl *template.Template, data TemplateData, locale, websiteDir string) (err error) {
	pi := path.Join(websiteDir, locale, "index.html")
	p4 := path.Join(websiteDir, locale, "404.html")
	data.OgLocale = locale

	println(pi)
	println(p4)
	println(data.OgLocale)

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
	pathconffile := flag.String("pathconf", "", "JSON config file for build path")
	flag.Parse()

	siteconf, err := LoadJsonConfig(*siteconffile)
	if err != nil {
		panic(err)
	}
	pathconf, err := LoadJsonConfig(*pathconffile)
	if err != nil {
		panic(err)
	}

	data := TemplateData{
		SiteUrl:     siteconf["SiteUrl"],
		TipitakaURL: siteconf["TipitakaURL"],
		OgImage:     siteconf["OgImage"],
		OgUrl:       siteconf["OgUrl"],
		OgLocale:    siteconf["OgLocale"],
	}

	tmpl, err := gtmpl.ParseDirectoryTree("messages", pathconf["localeDir"], pathconf["htmlTemplateDir"])
	if err != nil {
		panic(err)
	}

	err = CreateIndexAnd404(tmpl, data, "", pathconf["websiteDir"])

	var supportedLocales = []string{"en_US", "zh_TW", "vi_VN", "fr_FR"}
	for _, locale := range supportedLocales {
		gtmpl.SetLanguage(locale)
		err = CreateIndexAnd404(tmpl, data, locale, pathconf["websiteDir"])
		if err != nil {
			panic(err)
		}
	}
}
