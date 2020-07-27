package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
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
				ipp.SetValue(lib.RemoveLastChar(ipp.Value()))
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

func setupKeypad() {
	// pali virtual keypad
	bindKeypad("word", "keypad")

	// toggle virtual keypad
	tk := Document.GetElementById("toggle-keypad")
	kp := Document.GetElementById("keypad")
	tk.AddEventListener("click", func(e Event) {
		kp.ClassList().Toggle("is-hidden")

		spans := tk.QuerySelectorAll("span")
		for _, span := range spans {
			span.ClassList().Toggle("is-hidden")
		}
	})
}
