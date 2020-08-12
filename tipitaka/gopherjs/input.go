package main

import (
	"strings"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

var input *Object

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
func inputKeyupEventHandler(key string) {
	switch key {
	case "ArrowUp", "Up":
		println("ArrowUp")
	case "ArrowDown", "Down":
		println("ArrowDown")
	case "Enter":
		word := GetInputValue()
		if dicmgr.Lookup(word) {
			SetModalTitle(wordLinkHtml(word))
			go showWordDefinitionInModal(word)
		}
	default:
		word := GetInputValue()
		SetModalWords(GetSuggestedWordsHtml(word, 7))
	}
}

func FocusInput() {
	input.Focus()
}

func SetInputValue(v string) {
	input.SetValue(v)
}

func GetInputValue() string {
	s := input.Value()
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func SetupModalInput(selector string) {
	input = Document.QuerySelector(selector)
	input.AddEventListener("keyup", func(e Event) {
		inputKeyupEventHandler(e.Key())
	})
}
