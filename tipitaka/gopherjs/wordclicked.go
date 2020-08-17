package main

import (
	"bytes"
	"html/template"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
	"github.com/siongui/gopalilib/libfrontend"
)

func wordLinkHtml(word string) string {
	return "<a href='" + libfrontend.DictionarySuttaWordUrl(word) + "' target='_blank'>" + word + "</a>"
}

const pwt = `
<style>
.previewWordName > a {
  color: GoldenRod;
}

div.is-possible-word:hover {
  color: red;
  background-color: #F0F8FF;
  cursor: pointer;
}
</style>
{{range $i, $possibleWord := .}}
  <div class="is-possible-word is-size-5"
       onmouseenter="pwmeh({{$i}}, '{{$possibleWord}}')"
       onclick="pwh('{{$possibleWord}}')">
         {{$possibleWord}}
  </div>
{{end}}
`

func possibleWordClickHandler(word string) {
	SetModalContent("Loading " + wordLinkHtml(word) + " ...")

	go func() {
		wi, err := lib.HttpGetWordJson(libfrontend.HttpWordJsonPath(word))
		if err != nil {
			SetModalContent("Fail to Get " + word + ": " + err.Error())
			return
		}
		setting := lib.GetDefaultPaliSetting()

		html := `<div class="previewWordName is-size-4 mb-1">` + wordLinkHtml(word) + `</div>`
		html += dicmgr.GetWordDefinitionHtml(wi, setting, Window.Navigator().Languages())
		SetModalContent(html)
	}()
}

func possibleWordMouseenterHandler(i int, word string) {
	SetStateMachineCurrentIndexAndWord(i, word)
}

func GetSuggestedWordsHtml(word string, limit int) string {
	Document.Set("pwh", possibleWordClickHandler)
	Document.Set("pwmeh", possibleWordMouseenterHandler)

	t, err := template.New("pwt").Parse(pwt)
	if err != nil {
		return err.Error()
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, dicmgr.GetSuggestedWords(word, 7))
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func showWordDefinitionInModal(word string) {
	//showLookingUp()
	//defer hideLookingUp()
	wi, err := lib.HttpGetWordJson(libfrontend.HttpWordJsonPath(word))
	if err != nil {
		SetModalContent("Fail to Get " + word + ": " + err.Error())
		return
	}
	setting := lib.GetDefaultPaliSetting()
	SetModalContent(dicmgr.GetWordDefinitionHtml(wi, setting, Window.Navigator().Languages()))
}

func FindLongestPrefixWithNonZeroSuggestedWords(word string) string {
	for len(word) > 0 {
		if len(dicmgr.GetSuggestedWords(word, 10)) > 0 {
			break
		}
		word = lib.RemoveLastChar(word)
	}
	return word
}

func showPossibleWords(word string) {
	prefix := FindLongestPrefixWithNonZeroSuggestedWords(word)

	SetModalWords(GetSuggestedWordsHtml(prefix, 7))
	SetInputValue(prefix)
	ShowModalInput()
}

func wordClickedHandler(word string) {
	FocusInput()
	if dicmgr.Lookup(word) {
		SetModalTitle(wordLinkHtml(word))
		go showWordDefinitionInModal(word)
	} else {
		SetModalTitle(word)
		showPossibleWords(word)
	}
	openModal()
}
