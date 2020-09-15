package main

import (
	"strings"
	"time"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib/dicmgr"
	dic "github.com/siongui/gopalilib/lib/dictionary"
	"github.com/siongui/gopalilib/libfrontend/setting"
	"github.com/siongui/gopalilib/libfrontend/velthuis"
	sg "github.com/siongui/gopherjs-input-suggest"
	palitrans "github.com/siongui/pali-transliteration"
)

var mainContent *Object

func handleEnterEvent(input *Object) {
	w := strings.ToLower(strings.TrimSpace(input.Value()))
	input.Blur()
	go httpGetWordJson(w, true)
}

func toThai(input *Object) {
	w := strings.ToLower(strings.TrimSpace(input.Value()))
	t := palitrans.RomanToThai(w)
	Document.GetElementById("r2t").SetInnerHTML(t)
}

func handleInputKeyUp(e Event) {
	switch keycode := e.KeyCode(); keycode {
	case 13:
		// user press enter key
		handleEnterEvent(e.Target())
	default:
	}
}

func main() {
	setting.SetStorageKeyName("PaliDictionarySetting")
	setting.SetupPaliSetting()
	setupKeypad()

	// add pali input method to input text element
	// call velthuis before input suggest setup (order of keyevent handlers matters)
	velthuis.BindPaliInputMethodToInputTextElementById("word")

	// toggle type hint table
	tth := Document.GetElementById("toggle-type-hint")
	tht := Document.QuerySelector(".pali-type-hint-table")
	tth.AddEventListener("click", func(e Event) {
		tht.ClassList().Toggle("is-hidden")

		spans := tth.QuerySelectorAll("span")
		for _, span := range spans {
			span.ClassList().Toggle("is-hidden")
		}
	})

	// init variables
	mainContent = Document.GetElementById("main-content")

	// input suggest menu
	sg.BindSuggest("word", func(w string) []string {
		w = strings.ToLower(w)
		Document.GetElementById("word").SetValue(w)
		return dicmgr.GetSuggestedWords(w, 30)
	})
	// add Bulma css helper to input suggest menu
	ism := Document.QuerySelector(".suggest")
	ism.ClassList().Add("px-1")
	ism.ClassList().Add("py-1")
	ism.ClassList().Add("is-size-5")
	// setup word preview
	setupWordPreview()

	input := Document.GetElementById("word")
	input.AddEventListener("keyup", handleInputKeyUp)
	Document.AddEventListener("keyup", func(e Event) {
		// TAB: keyCode = 9
		if e.KeyCode() == 9 {
			if !input.IsFocused() {
				input.Focus()
			}
		}
	})

	// Hide loader and show input element while website is fully loaded.
	Window.AddEventListener("load", func(e Event) {
		si := Document.GetElementById("site-info")
		siteurl := si.Dataset().Get("siteurl").String()
		sitelocale := si.Dataset().Get("locale").String()
		dic.SetSiteUrl(siteurl)
		dic.SetCurrentLocale(sitelocale)

		setupContentAccordingToUrlPath()

		l := Document.GetElementById("website-loading")
		l.ClassList().Add("is-hidden")
		Document.QuerySelector("section.section").ClassList().Remove("is-hidden")

		setupBrowseDictionary()
	})

	// change url without reload
	Window.AddEventListener("popstate", func(e Event) {
		setupContentAccordingToUrlPath()
		/*
			if e.Get("state") == nil {
				// do nothing
			} else {
				setupContentAccordingToUrlPath()
				// state here stores pali word
				//word := e.Get("state").String()
				//go httpGetWordJson(word, false)
			}
		*/
	})

	// Romanized Pali to Thai
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				toThai(input)
			}
		}
	}()
}
