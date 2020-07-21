package main

import (
	"strings"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	jsgettext "github.com/siongui/gopherjs-i18n"
)

var supportedLocales = []string{"en_US", "zh_TW", "vi_VN", "fr_FR"}
var navigatorLanguages = Window.Navigator().Languages()

func setDocumentTitle(titleLocale string, typ lib.PageType, wordOrPrefix string) {
	//title := jsgettext.Gettext(titleLocale, "Pali Dictionary | Pāli to English, Chinese, Japanese, Vietnamese, Burmese Dictionary")
	title := jsgettext.Gettext(titleLocale, "Pāli Dictionary")
	if typ == lib.AboutPage {
		// add prefix "About"?
	}
	if typ == lib.WordPage {
		title = wordOrPrefix + " - " + jsgettext.Gettext(titleLocale, "Definition and Meaning") + " - " + title
	}
	if typ == lib.PrefixPage {
		title = jsgettext.Gettext(titleLocale, "Words Start with") + " " + wordOrPrefix + " - " + title
	}
	Document.Set("title", title)
}

func getFinalShowLocale() string {
	// show language according to site url and NavigatorLanguages API
	locale := Document.GetElementById("site-info").Dataset().Get("locale").String()
	if locale == "" {
		return jsgettext.DetermineLocaleByNavigatorLanguages(navigatorLanguages, supportedLocales)
	}
	return locale
}

func setupContentAccordingToUrlPath() {
	// show language according to NavigatorLanguages API
	titleLocale := getFinalShowLocale()
	jsgettext.Translate(titleLocale)

	up := Window.Location().Pathname()
	typ := lib.DeterminePageType(up)
	if typ == lib.RootPage {
		mainContent.RemoveAllChildNodes()
		setDocumentTitle(titleLocale, lib.RootPage, "")
		// maybe put some news in the future.
		return
	}
	if typ == lib.AboutPage {
		mainContent.RemoveAllChildNodes()
		mainContent.SetInnerHTML(Document.GetElementById("about").InnerHTML())
		setDocumentTitle(titleLocale, lib.AboutPage, "")
		return
	}
	if typ == lib.WordPage {
		mainContent.RemoveAllChildNodes()
		w := lib.GetWordFromUrlPath(up)
		setDocumentTitle(titleLocale, lib.WordPage, w)
		//println(w)
		go httpGetWordJson(w, false)
		return
	}
	if typ == lib.PrefixPage {
		mainContent.RemoveAllChildNodes()
		p := lib.GetPrefixFromUrlPath(up)
		setDocumentTitle(titleLocale, lib.PrefixPage, p)
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
