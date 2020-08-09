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
		if len(dicmgr.GetSuggestedWords(word, 10)) > 0 {
			break
		}
		word = lib.RemoveLastChar(word)
	}

	SetModalBody(GetPossibleWordsHtml(word, dicmgr.GetSuggestedWords(word, 7)))
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
