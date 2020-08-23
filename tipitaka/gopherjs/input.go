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
	CurrentWord  string
	Words        []string
}

func (s *StateMachine) HandleArrowUp() {
	if s.CurrentIndex >= 0 && s.CurrentIndex < len(s.Words) {
		UnhighlightWord(s.CurrentIndex, s.Words[s.CurrentIndex])
	}
	s.CurrentIndex--
	if s.CurrentIndex < 0 {
		s.CurrentIndex = 0
	}
	s.CurrentWord = s.Words[s.CurrentIndex]
	SetInputValue(s.CurrentWord)
	if s.CurrentIndex >= 0 && s.CurrentIndex < len(s.Words) {
		HighlightWord(s.CurrentIndex, s.Words[s.CurrentIndex])
	}
}

func (s *StateMachine) HandleArrowDown() {
	if s.CurrentIndex >= 0 && s.CurrentIndex < len(s.Words) {
		UnhighlightWord(s.CurrentIndex, s.Words[s.CurrentIndex])
	}
	s.CurrentIndex++
	if s.CurrentIndex == len(s.Words) {
		s.CurrentIndex = len(s.Words) - 1
	}
	s.CurrentWord = s.Words[s.CurrentIndex]
	SetInputValue(s.CurrentWord)
	if s.CurrentIndex >= 0 && s.CurrentIndex < len(s.Words) {
		HighlightWord(s.CurrentIndex, s.Words[s.CurrentIndex])
	}
}

func (s *StateMachine) HandleEnter() {
	word := GetInputValue()
	if dicmgr.Lookup(word) {
		SetModalTitle(wordLinkHtml(word))
		ResetStateMachine(word)
		go showWordDefinitionInModal(word)
	}
}

func SetStateMachineCurrentIndexAndWord(i int, word string) {
	st.CurrentIndex = i
	st.CurrentWord = word
	SetInputValue(word)
}

func (s *StateMachine) HandleDefault() {
	word := GetInputValue()
	ResetStateMachine(word)
}

func ResetStateMachine(word string) {
	st.CurrentIndex = -1
	st.CurrentWord = word
	st.Words = dicmgr.GetSuggestedWords(word, 7)
	SetModalWords(GetSuggestedWordsHtml(word, 7))
	SetInputValue(word)
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
