package main

import (
	"encoding/json"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/libfrontend"
	"github.com/siongui/gopalilib/libfrontend/everyword"
	"github.com/siongui/gopalilib/libfrontend/setting"
	"github.com/siongui/gopalilib/libfrontend/treeview"
	"github.com/siongui/gopalilib/libfrontend/velthuis"
	"github.com/siongui/gopalilib/libfrontend/xslt"
)

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
