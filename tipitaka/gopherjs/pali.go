package main

import (
	"encoding/json"
	"strings"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/libfrontend"
	"github.com/siongui/gopalilib/libfrontend/everyword"
	"github.com/siongui/gopalilib/libfrontend/setting"
	"github.com/siongui/gopalilib/libfrontend/treeview"
	"github.com/siongui/gopalilib/libfrontend/velthuis"
	"github.com/siongui/gopalilib/libfrontend/xslt"
	palitrans "github.com/siongui/pali-transliteration"
)

func toThai() {
	mainview := Document.GetElementById("mainview")
	spans := mainview.QuerySelectorAll("span[class=\"paliword\"]")
	for _, span := range spans {
		word := strings.ToLower(span.InnerHTML())
		span.SetInnerHTML(palitrans.RomanToThai(word))
	}
}

func createToThaiButton() *Object {
	r2t := Document.CreateElement("button")
	r2t.ClassList().Add("toThai")
	r2t.ClassList().Add("button")
	r2t.ClassList().Add("is-small")
	r2t.ClassList().Add("is-rounded")
	r2t.SetInnerHTML("To Thai Script (experimental)")
	r2t.AddEventListener("click", func(e Event) {
		r2t.ClassList().Add("is-rounded")
		toThai()
		r2t.ClassList().Add("is-hidden")
	})
	return r2t
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
	mainview.QuerySelector("div.content").Call("prepend", createToThaiButton())
	ToggleMobileTreeview()
}

func main() {
	libfrontend.TranslateDocument(libfrontend.GetFinalShowLocale())

	b, _ := ReadFile("tpktoc.json")
	//println(string(b))
	tree := lib.Tree{}
	json.Unmarshal(b, &tree)
	treeview.NewTreeview("treeview", tree, xmlAction)
	SetupTipitakaUrl(tree)

	xslt.SetupXSLTProcessor(libfrontend.GetXslUrl())
	SetupModal()
	SetupMobileTreeviewToggle()
	// Call velthuis before SetupModalInput (order of keyevent handler matters)
	velthuis.BindPaliInputMethodToInputTextElementById("modal-input")
	SetupModalInput("#modal-input")
	setting.SetupPaliSetting()

	HideIsLoadingWebsite()
}
