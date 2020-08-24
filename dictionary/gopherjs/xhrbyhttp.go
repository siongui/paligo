package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
	dic "github.com/siongui/gopalilib/lib/dictionary"
	"github.com/siongui/gopalilib/libfrontend"
	"github.com/siongui/gopalilib/libfrontend/setting"
)

func hideLookingUp() {
	l := Document.GetElementById("looking-up")
	l.ClassList().Add("is-hidden")
}

func showLookingUp() {
	l := Document.GetElementById("looking-up")
	l.ClassList().Remove("is-hidden")
}

func httpGetWordJson(w string, changeUrl bool) {
	showLookingUp()
	defer hideLookingUp()

	wi, err := lib.HttpGetWordJson(libfrontend.HttpWordJsonPath(w))
	if err != nil {
		mainContent.Set("textContent", err.Error())
		return
	}

	if changeUrl {
		Window.History().PushState(w, "", dic.WordUrlPath(w))
		setDocumentTitle(getFinalShowLocale(), dic.WordPage, w)
	}

	mainContent.RemoveAllChildNodes()
	mainContent.Set("innerHTML", dicmgr.GetWordDefinitionHtml(wi, setting.LoadPaliSetting(), Window.Navigator().Languages()))
}
