package main

import (
	"encoding/json"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	jsgettext "github.com/siongui/gopherjs-i18n"
	"github.com/siongui/paliDataVFS"
)

func getFinalShowLocale() string {
	var supportedLocales = []string{"en_US", "zh_TW", "vi_VN", "fr_FR"}
	var navigatorLanguages = Window.Navigator().Languages()
	// show language according to site url and NavigatorLanguages API
	locale := Document.GetElementById("site-info").Dataset().Get("locale").String()
	if locale == "" {
		return jsgettext.DetermineLocaleByNavigatorLanguages(navigatorLanguages, supportedLocales)
	}
	return locale
}

func main() {
	//println(getFinalShowLocale())
	jsgettext.SetupTranslationMapping(paliDataVFS.GetPoJsonBlob())
	jsgettext.Translate(getFinalShowLocale())

	b, _ := ReadFile("tpktoc.json")
	//println(string(b))
	tree := lib.Tree{}
	json.Unmarshal(b, &tree)
	NewTreeview("treeview", tree)
}
