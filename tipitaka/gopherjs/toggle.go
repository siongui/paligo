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

func ShowIsLoadingXML() {
	for _, l := range Document.QuerySelectorAll(".is-loading-xml") {
		l.ClassList().Remove("is-hidden")
	}
}

func HideIsLoadingXML() {
	for _, l := range Document.QuerySelectorAll(".is-loading-xml") {
		l.ClassList().Add("is-hidden")
	}
}

func HideIsLoadingTreeview() {
	Document.QuerySelector(".is-loading-treeview").ClassList().Add("is-hidden")
}
