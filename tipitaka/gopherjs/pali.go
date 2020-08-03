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

func xmlAction(action string) {
	mainview := Document.GetElementById("mainview")
	mainview.QuerySelector("div.notification").ClassList().Remove("is-hidden")

	// Load the xml file using synchronous (third param is set to false) XMLHttpRequest
	myXMLHTTPRequest := NewXMLHttpRequest()
	myXMLHTTPRequest.Open("GET", ActionXmlUrl(action), false)
	myXMLHTTPRequest.Send()

	xmlDoc := myXMLHTTPRequest.ResponseXML()
	fragment := GetXSLTProcessor().TransformToFragment(xmlDoc, Document)

	mainview.QuerySelector("div.content").RemoveAllChildNodes()
	mainview.QuerySelector("div.content").AppendChild(fragment)

	MarkEveryWord("#mainview > div.content", func(word string) {
		SetModalTitle(word)
		openModal()
	})

	mainview.QuerySelector("div.notification").ClassList().Add("is-hidden")
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

	SetupXSLTProcessor()
	SetupModal()

	Document.GetElementById("treeview").QuerySelector("div.notification").ClassList().Add("is-hidden")
}
