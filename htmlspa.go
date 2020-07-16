// Create the only index.html for Single Page Application (SPA), and 404.html to
// be served if not found.
package main

import (
	"encoding/json"
	"flag"
	"os"
	"path"

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

	findex, err := os.Create(path.Join(pathconf["websiteDir"], "index.html"))
	if err != nil {
		panic(err)
	}
	defer findex.Close()
	f404, err := os.Create(path.Join(pathconf["websiteDir"], "404.html"))
	if err != nil {
		panic(err)
	}
	defer f404.Close()

	tmpl, err := gtmpl.ParseDirectoryTree("messages", pathconf["localeDir"], pathconf["htmlTemplateDir"])
	if err != nil {
		panic(err)
	}

	err = tmpl.ExecuteTemplate(findex, "index.html", &data)
	if err != nil {
		panic(err)
	}
	err = tmpl.ExecuteTemplate(f404, "404.html", &data)
	if err != nil {
		panic(err)
	}
}
