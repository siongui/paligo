package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib/tipitaka"
)

func openTreeview(span *Object) {
	parent := span.ParentNode()
	if parent.ClassList().Contains("childrenContainer") {
		//println("childrenContainer")
		parent.Get("previousSibling").Get("lastChild").Call("click")
		openTreeview(parent)
	} else if parent.Get("id").String() == "treeview" {
		//println("end")
		return
	} else {
		//println("go up one level")
		openTreeview(parent)
	}
}

func checkIfPaliTextPage() {
	ok, paliTextPath := tipitaka.IsValidPaliTextUrlPath(Window.Location().Pathname())
	if !ok {
		return
	}

	//println(paliTextPath)
	span := Document.QuerySelector("[data-pali-text-path='" + paliTextPath + "']")
	// TODO FIXME: check if span is null here

	//println(span.InnerHTML())
	openTreeview(span)
}
