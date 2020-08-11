package main

import (
	. "github.com/siongui/godom"
)

// https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent/key
func modalInputKeyupEventHandler(key string) {
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

func SetupModalInput(selector string) {
	input := Document.QuerySelector(selector)
	input.AddEventListener("keyup", func(e Event) {
		modalInputKeyupEventHandler(e.Key())
	})
}
