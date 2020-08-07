package main

import (
	"net/http"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

func showWordDefinitionInModal(word string) {
	//showLookingUp()
	//defer hideLookingUp()

	resp, err := http.Get(HttpWordJsonPath(word))
	if err != nil {
		SetModalBody("Fail to look up " + word + err.Error())
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		SetModalBody("Fail to look up " + word + " != 200")
		return
	}

	wi, err := lib.DecodeHttpRespWord(resp.Body)
	if err != nil {
		SetModalBody("Fail to look up " + word + err.Error())
		return
	}

	setting := lib.PaliSetting{
		IsShowWordPreview: false,
		P2en:              true,
		P2ja:              true,
		P2zh:              true,
		P2vi:              true,
		P2my:              true,
		DicLangOrder:      "hdr",
	}

	SetModalBody(dicmgr.GetWordDefinitionHtml(wi, setting, Window.Navigator().Languages()))
}

func wordClickedHandler(word string) {
	SetModalTitle(word)
	if dicmgr.Lookup(word) {
		go showWordDefinitionInModal(word)
	} else {
		SetModalBody("not in trie")
	}
	openModal()
}
