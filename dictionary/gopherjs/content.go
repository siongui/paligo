package main

import (
	"strings"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib/dicmgr"
	dic "github.com/siongui/gopalilib/lib/dictionary"
	"github.com/siongui/gopalilib/lib/jsgettext"
	"github.com/siongui/gopalilib/libfrontend"
)

func setDocumentTitle(titleLocale string, typ dic.PageType, wordOrPrefix string) {
	//title := jsgettext.Gettext(titleLocale, "Pali Dictionary | Pāli to English, Chinese, Japanese, Vietnamese, Burmese Dictionary")
	title := jsgettext.Gettext(titleLocale, "Pāli Dictionary")
	if typ == dic.AboutPage {
		// add prefix "About"?
	}
	if typ == dic.WordPage {
		title = wordOrPrefix + " - " + jsgettext.Gettext(titleLocale, "Definition and Meaning") + " - " + title
	}
	if typ == dic.PrefixPage {
		title = jsgettext.Gettext(titleLocale, "Words Start with") + " " + wordOrPrefix + " - " + title
	}
	Document.Set("title", title)
}

func setupContentAccordingToUrlPath() {
	// show language according to NavigatorLanguages API
	titleLocale := libfrontend.GetFinalShowLocale()
	libfrontend.TranslateDocument(titleLocale)

	up := Window.Location().Pathname()
	typ := dic.DeterminePageType(up)
	if typ == dic.RootPage {
		mainContent.RemoveAllChildNodes()
		setDocumentTitle(titleLocale, dic.RootPage, "")
		// maybe put some news in the future.
		return
	}
	if typ == dic.AboutPage {
		mainContent.RemoveAllChildNodes()
		mainContent.SetInnerHTML(Document.GetElementById("about").InnerHTML())
		setDocumentTitle(titleLocale, dic.AboutPage, "")
		return
	}
	if typ == dic.WordPage {
		mainContent.RemoveAllChildNodes()
		w := dic.GetWordFromUrlPath(up)
		setDocumentTitle(titleLocale, dic.WordPage, w)
		//println(w)
		go httpGetWordJson(w, false)
		return
	}
	if typ == dic.PrefixPage {
		mainContent.RemoveAllChildNodes()
		p := dic.GetPrefixFromUrlPath(up)
		setDocumentTitle(titleLocale, dic.PrefixPage, p)
		//mainContent.SetInnerHTML("prefix " + p)
		prefixwords := dicmgr.GetSuggestedWords(p, 1000000)
		html := ""
		for _, prefixword := range prefixwords {
			html += `<li><a href="` + dic.WordUrlPath(prefixword) + `">` + prefixword + `</a></li>`
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
		html := `<li><a href="` + dic.PrefixUrlPath(prefix) + `">{{PREFIX}}</a></li>`
		html = strings.Replace(html, "{{PREFIX}}", prefix, 1)
		all += html
	}
	pl.SetInnerHTML(all)
}
