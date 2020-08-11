package main

import (
	. "github.com/siongui/godom"
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
		println("Enter")
	default:
		println("default")
	}
}

func SetInputValue(v string) {
	input.SetValue(v)
}

func SetupModalInput(selector string) {
	input = Document.QuerySelector(selector)
	input.AddEventListener("keyup", func(e Event) {
		inputKeyupEventHandler(e.Key())
	})
}
