// Create the only HTML for Single Page Application (SPA) for Pali dictionary
// website.
package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/siongui/gopalilib/dicutil"
)

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

	data := dicutil.TemplateData{
		SiteUrl:     siteconf["SiteUrl"],
		TipitakaURL: siteconf["TipitakaURL"],
		OgImage:     siteconf["OgImage"],
		OgUrl:       siteconf["OgUrl"],
		OgLocale:    siteconf["OgLocale"],
	}

	err = dicutil.CreateHTML(os.Stdout, "index.html", &data, pathconf["localeDir"], pathconf["htmlTemplateDir"])
	if err != nil {
		panic(err)
	}
}
