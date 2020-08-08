package main

import (
	. "github.com/siongui/godom"
)

func SetupMobileTreeviewToggle() {
	tv := Document.GetElementById("treeview")
	mtt := Document.GetElementById("mobile-treeview-toggle")
	mtt.AddEventListener("click", func(e Event) {
		tv.ClassList().Toggle("is-hidden-mobile")
		mtt.ClassList().Toggle("is-pulled-right")
		for _, span := range mtt.QuerySelectorAll("span") {
			span.ClassList().Toggle("is-hidden")
		}
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
