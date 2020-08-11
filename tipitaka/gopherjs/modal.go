package main

import (
	. "github.com/siongui/godom"
)

var modal *Object

func SetModalTitle(title string) {
	//modal.QuerySelector(".modal-card-title").SetInnerHTML(title)
	modal.QuerySelector("#modal-title").SetInnerHTML(title)
}

func SetModalBody(b string) {
	modal.QuerySelector(".modal-card-body").SetInnerHTML(b)
}

func openModal() {
	Document.DocumentElement().ClassList().Add("is-clipped")
	modal.ClassList().Add("is-active")
}

func closeModal() {
	Document.DocumentElement().ClassList().Remove("is-clipped")
	modal.ClassList().Remove("is-active")

	//SetModalBody("")
	SetModalWords("")
	SetModalContent("")
	SetInputValue("")
	HideModalInput()
}

func SetModalContent(html string) {
	Document.GetElementById("modal-content").SetInnerHTML(html)
}

func SetModalWords(html string) {
	Document.GetElementById("words").SetInnerHTML(html)
}

func HideModalInput() {
	Document.GetElementById("modal-input-toggle").Set("checked", false)
}

func ShowModalInput() {
	Document.GetElementById("modal-input-toggle").Set("checked", true)
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
