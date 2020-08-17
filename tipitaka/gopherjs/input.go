package main

import (
	"strings"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

var st *StateMachine

type StateMachine struct {
	Input        *Object
	CurrentIndex int
}

func (s *StateMachine) HandleArrowUp() {
	println("ArrowUp")
}

func (s *StateMachine) HandleArrowDown() {
	println("ArrowDown")
}

func (s *StateMachine) HandleEnter() {
	word := GetInputValue()
	if dicmgr.Lookup(word) {
		SetModalTitle(wordLinkHtml(word))
		go showWordDefinitionInModal(word)
	}
}

func (s *StateMachine) HandleDefault() {
	word := GetInputValue()
	SetModalWords(GetSuggestedWordsHtml(word, 7))
}

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
func inputKeyupEventHandler(key string) {
	switch key {
	case "ArrowUp", "Up":
		st.HandleArrowUp()
	case "ArrowDown", "Down":
		st.HandleArrowDown()
	case "Enter":
		st.HandleEnter()
	default:
		st.HandleDefault()
	}
}

func FocusInput() {
	st.Input.Focus()
}

func SetInputValue(v string) {
	st.Input.SetValue(v)
}

func GetInputValue() string {
	s := st.Input.Value()
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	return s
}

func SetupModalInput(selector string) {
	st = &StateMachine{}
	st.Input = Document.QuerySelector(selector)
	st.Input.AddEventListener("keyup", func(e Event) {
		inputKeyupEventHandler(e.Key())
	})
}
