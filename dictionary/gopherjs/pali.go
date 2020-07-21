package main

import (
	"strings"

	imepali "github.com/siongui/go-online-input-method-pali"
	bits "github.com/siongui/go-succinct-data-structure-trie"
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"

	sg "github.com/siongui/gopherjs-input-suggest"
	"github.com/siongui/paliDataVFS"
)

var mainContent *Object
var bookIdAndInfos = paliDataVFS.GetBookIdAndInfos()
var frozenTrie bits.FrozenTrie

func handleEnterEvent(input *Object) {
	w := strings.ToLower(strings.TrimSpace(input.Value()))
	input.Blur()
	go httpGetWordJson(w, true)
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
	// add pali input method to input text element
	imepali.BindPaliInputMethodToInputTextElementById("word")

	// pali virtual keypad
	bindKeypad("word", "keypad")

	// toggle virtual keypad
	tk := Document.GetElementById("toggle-keypad")
	kp := Document.GetElementById("keypad")
	tk.AddEventListener("click", func(e Event) {
		kp.ClassList().Toggle("is-hidden")

		spans := tk.QuerySelectorAll("span")
		for _, span := range spans {
			span.ClassList().Toggle("is-hidden")
		}
	})

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

	// init trie for words suggestion
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	frozenTrie = bits.FrozenTrie{}
	frozenTrie.Init(paliDataVFS.GetTrieData())

	// input suggest menu
	sg.BindSuggest("word", func(w string) []string {
		w = strings.ToLower(w)
		Document.GetElementById("word").SetValue(w)
		return frozenTrie.GetSuggestedWords(w, 30)
	})
	// add Bulma css helper to input suggest menu
	ism := Document.QuerySelector(".suggest")
	ism.ClassList().Add("px-1")
	ism.ClassList().Add("py-1")
	ism.ClassList().Add("is-size-5")

	setupNavbar()
	setupSetting()

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
		lib.SetSiteUrl(siteurl)
		lib.SetCurrentLocale(sitelocale)

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
}
