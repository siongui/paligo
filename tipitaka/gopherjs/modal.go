package main

import (
	. "github.com/siongui/godom"
)

var modal *Object

func SetModalTitle(title string) {
	modal.QuerySelector(".modal-card-title").SetInnerHTML(title)
}

func openModal() {
	Document.DocumentElement().ClassList().Add("is-clipped")
	modal.ClassList().Add("is-active")
}

func closeModal() {
	Document.DocumentElement().ClassList().Remove("is-clipped")
	modal.ClassList().Remove("is-active")
}

func SetupModal() {
	modal = Document.QuerySelector("div.modal")

	var closeElm []*Object
	closeElm = append(closeElm, modal.QuerySelector(".modal-background"))
	//closeElm = append(closeElm, modal.QuerySelector(".modal-close"))
	closeElm = append(closeElm, modal.QuerySelector(".modal-card-head .delete"))
	//closeElm = append(closeElm, modal.QuerySelector(".modal-card-foot .button"))
	for _, elm := range closeElm {
		elm.AddEventListener("click", func(e Event) {
			closeModal()
		})
	}

	Document.AddEventListener("keydown", func(e Event) {
		if e.KeyCode() == 27 {
			// ESC pressed
			closeModal()
		}
	})
}
