package main

import (
	. "github.com/siongui/godom"
)

//DISCUSS: close mobile nav menu after click?
func setupNavbar() {
	// setting nav item
	sl := Document.QuerySelector(".setting-link")
	sl.AddEventListener("click", func(e Event) {
		// toggle arrow
		downArrow := sl.QuerySelector(".down-arrow")
		downArrow.ClassList().Toggle("is-hidden")
		// right arrow
		downArrow.NextSibling().ClassList().Toggle("is-hidden")
		// setting menu
		Document.QuerySelector(".setting-menu").ClassList().Toggle("is-hidden")
	})
}
