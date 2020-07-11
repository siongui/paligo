package main

import (
	"strings"

	imepali "github.com/siongui/go-online-input-method-pali"
	bits "github.com/siongui/go-succinct-data-structure-trie"
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	jsgettext "github.com/siongui/gopherjs-i18n"
	sg "github.com/siongui/gopherjs-input-suggest"
	"github.com/siongui/paliDataVFS"
)

var mainContent *Object
var bookIdAndInfos = paliDataVFS.GetBookIdAndInfos()
var frozenTrie bits.FrozenTrie
var navigatorLanguages = Window.Navigator().Languages()

func isDev() bool {
	return Window.Location().Hostname() == "localhost"
}

func handleInputKeyUp(e Event) {
	switch keycode := e.KeyCode(); keycode {
	case 13:
		// user press enter key
		raw := e.Target().Value()
		raw = strings.TrimSpace(raw)
		w := strings.ToLower(raw)
		e.Target().Blur()
		go httpGetWordJson(w, true)
	default:
	}
}

func setupMainContentAccordingToUrlPath() {
	typ := lib.DeterminePageType(Window.Location().Pathname())
	if typ == lib.RootPage {
		mainContent.RemoveAllChildNodes()
	}
	if typ == lib.AboutPage {
		mainContent.RemoveAllChildNodes()
		mainContent.SetInnerHTML(Document.GetElementById("about").InnerHTML())
	}
	// TODO: handle other type of pages
}

func main() {
	// add pali input method to input text element
	imepali.BindPaliInputMethodToInputTextElementById("word")

	// init variables
	mainContent = Document.GetElementById("main-content")

	// init trie for words suggestion
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")
	frozenTrie = bits.FrozenTrie{}
	frozenTrie.Init(paliDataVFS.GetTrieData())

	// input suggest menu
	sg.BindSuggest("word", func(w string) []string {
		return frozenTrie.GetSuggestedWords(w, 30)
	})
	// add Bulma css helper to input suggest menu
	ism := Document.QuerySelector(".suggest")
	ism.ClassList().Add("px-1")
	ism.ClassList().Add("py-1")
	ism.ClassList().Add("is-size-5")

	setupNavbar()
	setupSetting()

	// show language according to NavigatorLanguages API
	supportedLocales := []string{"en_US", "zh_TW", "vi_VN", "fr_FR"}
	initialLocale := jsgettext.DetermineLocaleByNavigatorLanguages(navigatorLanguages, supportedLocales)
	if initialLocale != "en_US" {
		jsgettext.Translate(initialLocale)
	}

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
		l := Document.QuerySelector(".notification.is-info")
		l.ClassList().Add("is-hidden")
		input.ClassList().Remove("is-hidden")
	})

	// change url without reload
	Window.AddEventListener("popstate", func(e Event) {
		if e.Get("state") == nil {
			// do nothing
		} else {
			// state here stores pali word
			word := e.Get("state").String()
			go httpGetWordJson(word, false)
		}
	})

	setupMainContentAccordingToUrlPath()
}
