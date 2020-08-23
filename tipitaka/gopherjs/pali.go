package main

import (
	"encoding/json"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/jsgettext"
	"github.com/siongui/gopalilib/libfrontend"
	"github.com/siongui/gopalilib/libfrontend/everyword"
	"github.com/siongui/gopalilib/libfrontend/treeview"
	"github.com/siongui/gopalilib/libfrontend/velthuis"
	"github.com/siongui/gopalilib/libfrontend/xslt"
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

func xmlAction(t lib.Tree) {
	// FIXME: show loading not working on Chromium
	ShowIsLoadingXML(t.Text)
	defer HideIsLoadingXML()

	mainview := Document.GetElementById("mainview")

	// Load the xml file using synchronous (third param is set to false) XMLHttpRequest
	myXMLHTTPRequest := NewXMLHttpRequest()
	myXMLHTTPRequest.Open("GET", libfrontend.ActionXmlUrl(t.Action), false)
	myXMLHTTPRequest.Send()

	xmlDoc := myXMLHTTPRequest.ResponseXML()
	fragment := xslt.GetXSLTProcessor().TransformToFragment(xmlDoc, Document)

	mainview.QuerySelector("div.content").RemoveAllChildNodes()
	mainview.QuerySelector("div.content").AppendChild(fragment)

	everyword.MarkEveryWord("#mainview > div.content", wordClickedHandler)

	ToggleMobileTreeview()
}

func TranslateDocument(locale string) {
	elms := Document.QuerySelectorAll("[data-default-string]")
	for _, elm := range elms {
		str := elm.Get("dataset").Get("defaultString").String()
		elm.Set("textContent", jsgettext.Gettext(locale, str))
	}
}

func main() {
	//println(getFinalShowLocale())
	TranslateDocument(getFinalShowLocale())

	b, _ := ReadFile("tpktoc.json")
	//println(string(b))
	tree := lib.Tree{}
	json.Unmarshal(b, &tree)
	treeview.NewTreeview("treeview", tree, xmlAction)

	xslt.SetupXSLTProcessor(libfrontend.GetXslUrl())
	SetupModal()
	SetupMobileTreeviewToggle()
	// Call velthuis before SetupModalInput (order of keyevent handler matters)
	velthuis.BindPaliInputMethodToInputTextElementById("modal-input")
	SetupModalInput("#modal-input")

	HideIsLoadingWebsite()
}
