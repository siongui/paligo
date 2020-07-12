package main

import (
	"net/url"
	"strings"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
)

func setupMainContentAccordingToUrlPath() {
	up, _ := url.PathUnescape(Window.Location().Pathname())
	typ := lib.DeterminePageType(up)
	if typ == lib.RootPage {
		mainContent.RemoveAllChildNodes()
		// maybe put some news in the future.
		return
	}
	if typ == lib.AboutPage {
		mainContent.RemoveAllChildNodes()
		mainContent.SetInnerHTML(Document.GetElementById("about").InnerHTML())
		return
	}
	if typ == lib.WordPage {
		mainContent.RemoveAllChildNodes()
		w := lib.GetWordFromUrlPath(up)
		//println(w)
		go httpGetWordJson(w, false)
		return
	}
	if typ == lib.PrefixPage {
		mainContent.RemoveAllChildNodes()
		p := lib.GetPrefixFromUrlPath(up)
		//mainContent.SetInnerHTML("prefix " + p)
		prefixwords := frozenTrie.GetSuggestedWords(p, 1000000)
		html := ""
		for _, prefixword := range prefixwords {
			html += `<li><a href="` + lib.WordUrlPath(prefixword) + `">` + prefixword + `</a></li>`
		}
		mainContent.SetInnerHTML(`<nav class="breadcrumb" aria-label="breadcrumbs"><ul>` + html + `</ul></nav>`)
		return
	}
	// handle other type of pages?
}

func setupBrowseDictionary() {
	pl := Document.GetElementById("prefixList")
	prefixs := []string{"a", "ā", "b", "c", "d", "ḍ", "e", "g", "h", "i", "ī", "j", "k", "l", "ḷ", "m", "ŋ", "n", "ñ", "ṅ", "ṇ", "o", "p", "r", "s", "t", "ṭ", "u", "ū", "v", "y", "-", "°"}
	all := ""
	for _, prefix := range prefixs {
		html := `<li><a href="` + lib.PrefixUrlPath(prefix) + `">{{PREFIX}}</a></li>`
		html = strings.Replace(html, "{{PREFIX}}", prefix, 1)
		all += html
	}
	pl.SetInnerHTML(all)
}

func isDev() bool {
	return Window.Location().Hostname() == "localhost"
}
