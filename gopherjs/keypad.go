package main

import (
	"unicode/utf8"

	. "github.com/siongui/godom"
	sg "github.com/siongui/gopherjs-input-suggest"
)

var letters = [][]string{
	[]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	[]string{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
	[]string{"z", "x", "c", "v", "b", "n", "m"},
	[]string{"ā", "ḍ", "ī", "ḷ", "ṁ", "ṃ", "ñ", "ṇ", "ṭ", "ū", "ŋ", "ṅ"},
}

// http://blog.elliottcable.name/posts/useful_unicode.xhtml
// https://www.compart.com/en/unicode/U+23CE
var specialKeys = []string{"Backspace ⌫", "Enter ⏎"}

func RemoveLastChar(str string) string {
	for len(str) > 0 {
		_, size := utf8.DecodeLastRuneInString(str)
		return str[:len(str)-size]
	}
	return str
}

func initKeypadInputElement(value string) *Object {
	ie := Document.CreateElement("input")
	ie.Set("type", "button")
	ie.ClassList().Add("button")
	ie.ClassList().Add("is-small")
	ie.SetValue(value)
	return ie
}

// initialization function
func bindKeypad(inputId, keypadId string) {
	ipp := Document.GetElementById(inputId)
	k := Document.GetElementById(keypadId)

	for row := 0; row < len(letters); row++ {
		de := Document.CreateElement("div")
		for i := 0; i < len(letters[row]); i++ {
			ie := initKeypadInputElement(letters[row][i])
			ie.AddEventListener("click", func(e Event) {
				ipp.SetValue(ipp.Value() + ie.Value())
				sg.UpdateSuggestion()
			})
			de.AppendChild(ie)
		}
		k.AppendChild(de)
	}

	de2 := Document.CreateElement("div")
	for i := 0; i < len(specialKeys); i++ {
		ie := initKeypadInputElement(specialKeys[i])
		if i == 0 {
			// Backspace
			ie.AddEventListener("click", func(e Event) {
				ipp.SetValue(RemoveLastChar(ipp.Value()))
				sg.UpdateSuggestion()
			})
		}
		if i == 1 {
			// Enter
			ie.AddEventListener("click", func(e Event) {
				//handleEnterEvent(ipp)
				//sg.HideSuggestion()
				option := Window.Get("Object").New()
				option.Set("keyCode", 13)
				ke := Window.Get("KeyboardEvent").New("keyup", option)
				ipp.Call("dispatchEvent", ke)
			})
		}
		de2.AppendChild(ie)
	}
	k.AppendChild(de2)
}
