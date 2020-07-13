package main

import (
	. "github.com/siongui/godom"
	sg "github.com/siongui/gopherjs-input-suggest"
)

var letters = [][]string{
	[]string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p"},
	[]string{"a", "s", "d", "f", "g", "h", "j", "k", "l"},
	[]string{"z", "x", "c", "v", "b", "n", "m"},
	[]string{"ā", "ḍ", "ī", "ḷ", "ṁ", "ṃ", "ñ", "ṇ", "ṭ", "ū", "ŋ", "ṅ"},
}

// initialization function
func bindKeypad(inputId, keypadId string) {
	ipp := Document.GetElementById(inputId)
	k := Document.GetElementById(keypadId)

	for row := 0; row < len(letters); row++ {
		de := Document.CreateElement("div")
		for i := 0; i < len(letters[row]); i++ {
			ie := Document.CreateElement("input")
			ie.Set("type", "button")
			ie.ClassList().Add("button")
			ie.ClassList().Add("is-small")
			ie.SetValue(letters[row][i])
			ie.AddEventListener("click", func(e Event) {
				ori := ipp.Value()
				ipp.SetValue(ori + ie.Value())
				sg.UpdateSuggestion()
			})
			de.AppendChild(ie)
		}
		k.AppendChild(de)
	}
}
