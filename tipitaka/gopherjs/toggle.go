package main

import (
	. "github.com/siongui/godom"
)

func SetupMobileTreeviewToggle() {
	Document.GetElementById("mobile-treeview-toggle").AddEventListener("click", func(e Event) {
		ToggleMobileTreeview()
	})
}

func ToggleMobileTreeview() {
	tv := Document.GetElementById("treeview")
	mtt := Document.GetElementById("mobile-treeview-toggle")

	tv.ClassList().Toggle("is-hidden-mobile")
	mtt.ClassList().Toggle("is-pulled-right")
	for _, span := range mtt.QuerySelectorAll("span") {
		span.ClassList().Toggle("is-hidden")
	}
}

func ShowIsLoadingXML(text string) {
	for _, l := range Document.QuerySelectorAll(".is-loading-xml") {
		l.SetTextContent("Loading XML " + text + " ...")
		l.ClassList().Remove("is-hidden")
	}
}

func HideIsLoadingXML() {
	for _, l := range Document.QuerySelectorAll(".is-loading-xml") {
		l.ClassList().Add("is-hidden")
	}
}

func HideIsLoadingWebsite() {
	Document.QuerySelector(".is-loading-website").ClassList().Add("is-hidden")
	// show website content
	Document.QuerySelector("section.pt-1").ClassList().Remove("is-hidden")
}
