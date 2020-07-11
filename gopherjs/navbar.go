package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopherjs-i18n"
	"github.com/siongui/paliDataVFS"
)

//DISCUSS: close mobile nav menu after click?
func setupNavbar() {
	jsgettext.SetupTranslationMapping(paliDataVFS.GetPoJsonBlob())

	// about nav item
	al := Document.QuerySelector(".about-link")
	al.AddEventListener("click", func(e Event) {
		// load about content
		mainContent.RemoveAllChildNodes()
		mainContent.SetInnerHTML(Document.GetElementById("about").InnerHTML())
	})

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

	// language select nav item
	lss := Document.QuerySelectorAll(".lang-select")
	for _, ls := range lss {
		ls.AddEventListener("click", func(e Event) {
			// Cannot use the following line:
			//locale := ls.Dataset().Get("lang").String()
			// otherwise the lang value of dataset will always be the value of last "ls"
			// must replace "ls" with e.Target()
			locale := e.Target().Dataset().Get("lang").String()
			jsgettext.Translate(locale)
		})
	}

	// mobile toggle
	Document.AddEventListener("DOMContentLoaded", func(e Event) {
		nbs := Document.QuerySelectorAll(".navbar-burger")
		for _, nb := range nbs {
			nb.AddEventListener("click", func(e Event) {
				tg := e.Target().Dataset().Get("target").String()
				tgEl := Document.GetElementById(tg)

				e.Target().ClassList().Toggle("is-active")
				tgEl.ClassList().Toggle("is-active")
			})
		}
	})
}
