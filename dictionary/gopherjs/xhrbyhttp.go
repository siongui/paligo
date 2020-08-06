package main

import (
	"net/http"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	dic "github.com/siongui/gopalilib/lib/dictionary"
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
	resp, err := http.Get(HttpWordJsonPath(w))
	if err != nil {
		mainContent.Set("textContent", "Not Found")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		mainContent.Set("textContent", "Not Found")
		return
	}

	wi, err := lib.DecodeHttpRespWord(resp.Body)
	if err != nil {
		mainContent.Set("textContent", "Not Found")
		return
	}

	if changeUrl {
		Window.History().PushState(w, "", dic.WordUrlPath(w))
		setDocumentTitle(getFinalShowLocale(), dic.WordPage, w)
	}

	showWordByTemplate(wi)
}
