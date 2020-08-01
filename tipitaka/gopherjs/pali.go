package main

import (
	"encoding/json"
	"strings"

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

func xmlAction(action string) {
	url := ActionXmlUrl(action)

	xsltProcessor := NewXSLTProcessor()

	// Load the xsl file using synchronous (third param is set to false) XMLHttpRequest
	myXMLHTTPRequest := NewXMLHttpRequest()
	myXMLHTTPRequest.Open("GET", url, false)
	myXMLHTTPRequest.Send()

	xslRef := myXMLHTTPRequest.ResponseXML()

	// Finally import the .xsl
	xsltProcessor.ImportStylesheet(xslRef)

	// Cannot append DOM element to DIV node: Uncaught HierarchyRequestError: Failed to execute 'appendChild' on 'Node'
	// https://stackoverflow.com/a/29643573
	Document.GetElementById("mainview").RemoveAllChildNodes()
	content := xslRef.DocumentElement().QuerySelector("body").InnerHTML()
	Document.GetElementById("mainview").SetInnerHTML(strings.Replace(content, "rend", "class", -1))
}

func main() {
	//println(getFinalShowLocale())
	jsgettext.SetupTranslationMapping(paliDataVFS.GetPoJsonBlob())
	jsgettext.Translate(getFinalShowLocale())

	b, _ := ReadFile("tpktoc.json")
	//println(string(b))
	tree := lib.Tree{}
	json.Unmarshal(b, &tree)
	NewTreeview("treeview", tree, xmlAction)
}
