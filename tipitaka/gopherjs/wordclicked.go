package main

import (
	"fmt"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

func wordLinkHtml(word string) string {
	return fmt.Sprintf("<a href='%s' target='_blank'>%s</a>", wordDictionaryUrl(word), word)
}

func showWordDefinitionInModal(word string) {
	//showLookingUp()
	//defer hideLookingUp()
	wi, err := lib.HttpGetWordJson(HttpWordJsonPath(word))
	if err != nil {
		SetModalBody(fmt.Sprintf("Fail to Get %s: %s", word, err.Error()))
		return
	}
	setting := lib.GetDefaultPaliSetting()
	SetModalBody(dicmgr.GetWordDefinitionHtml(wi, setting, Window.Navigator().Languages()))
}

func showPossibleWords(word string) {
	for len(word) > 0 {
		word = lib.RemoveLastChar(word)
		if len(dicmgr.GetSuggestedWords(word, 10)) > 0 {
			break
		}
	}

	html := ""
	for _, w := range dicmgr.GetSuggestedWords(word, 10) {
		html += fmt.Sprintf("<div>%s</div>", wordLinkHtml(w))
	}

	SetModalBody(html)
}

func wordClickedHandler(word string) {
	SetModalTitle(word)
	if dicmgr.Lookup(word) {
		SetModalTitle(wordLinkHtml(word))
		go showWordDefinitionInModal(word)
	} else {
		showPossibleWords(word)
	}
	openModal()
}
