package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
	"github.com/siongui/gopalilib/libfrontend"
	"github.com/siongui/gopalilib/libfrontend/setting"
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

func httpGetWordJson2(word string) {
	wi, err := lib.HttpGetWordJson(libfrontend.HttpWordJsonPath(word))
	if err != nil {
		// TODO: handle error here.
		return
	}

	setWordPreviewUI(word, dicmgr.GetWordPreviewHtml(word, wi, setting.LoadPaliSetting(), Window.Navigator().Languages()))
}

func setupWordPreview() {
	wordPreviewElm = Document.QuerySelector(".suggestedWordPreview")
	sg.OnHighlightSelectedWord(func(word string) {
		//println(word)
		if !setting.LoadPaliSetting().IsShowWordPreview {
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
