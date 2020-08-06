package main

import (
	"net/http"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
	sg "github.com/siongui/gopherjs-input-suggest"
)

var wordPreviewElm *Object

func setWordPreviewUI(word, rawhtml string) {
	wordPreviewElm.SetInnerHTML(rawhtml)
	wordPreviewElm.ClassList().Remove("is-hidden")
	Document.QuerySelector(".suggest").ClassList().Add("suggest-is-absolute")
	w := Document.QuerySelector(".suggest").Get("offsetWidth").String() + "px"
	//println(w)
	wordPreviewElm.Style().SetLeft(w)
}

func showWordPreviewByTemplate(word string, wi lib.BookIdWordExps) {
	setWordPreviewUI(word, dicmgr.GetWordPreviewHtml(word, wi, getSetting(), navigatorLanguages))
}

func httpGetWordJson2(word string) {
	resp, err := http.Get(HttpWordJsonPath(word))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return
	}

	wi, err := lib.DecodeHttpRespWord(resp.Body)
	if err != nil {
		return
	}

	showWordPreviewByTemplate(word, wi)
}

func setupWordPreview() {
	wordPreviewElm = Document.QuerySelector(".suggestedWordPreview")
	sg.OnHighlightSelectedWord(func(word string) {
		//println(word)
		if !getSetting().IsShowWordPreview {
			return
		}
		//println("show word preview")
		go httpGetWordJson2(word)
	})

	sg.OnUpdateSuggestMenu(func(word string) {
		wordPreviewElm.ClassList().Add("is-hidden")
		Document.QuerySelector(".suggest").ClassList().Remove("suggest-is-absolute")
	})

	sg.OnHideSuggestMenu(func() {
		wordPreviewElm.ClassList().Add("is-hidden")
		Document.QuerySelector(".suggest").ClassList().Remove("suggest-is-absolute")
	})
}
