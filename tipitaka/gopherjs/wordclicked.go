package main

import (
	"bytes"
	"html/template"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

func wordLinkHtml(word string) string {
	return "<a href='" + wordDictionaryUrl(word) + "' target='_blank'>" + word + "</a>"
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
{{range $possibleWord := .}}
  <div class="is-possible-word is-size-5"
       onclick="pwh('{{$possibleWord}}')">
         {{$possibleWord}}
  </div>
{{end}}
`

func onPossibleWordHandler(word string) {
	SetModalContent("Loading " + wordLinkHtml(word) + " ...")

	go func() {
		wi, err := lib.HttpGetWordJson(HttpWordJsonPath(word))
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

func GetPossibleWordsHtml(word string, possibleWords []string) string {
	Document.Set("pwh", onPossibleWordHandler)

	t, err := template.New("pwt").Parse(pwt)
	if err != nil {
		return err.Error()
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, possibleWords)
	if err != nil {
		return err.Error()
	}
	return buf.String()
}

func showWordDefinitionInModal(word string) {
	//showLookingUp()
	//defer hideLookingUp()
	wi, err := lib.HttpGetWordJson(HttpWordJsonPath(word))
	if err != nil {
		SetModalContent("Fail to Get " + word + ": " + err.Error())
		return
	}
	setting := lib.GetDefaultPaliSetting()
	SetModalContent(dicmgr.GetWordDefinitionHtml(wi, setting, Window.Navigator().Languages()))
}

func showPossibleWords(word string) {
	for len(word) > 0 {
		if len(dicmgr.GetSuggestedWords(word, 10)) > 0 {
			break
		}
		word = lib.RemoveLastChar(word)
	}

	SetModalWords(GetPossibleWordsHtml(word, dicmgr.GetSuggestedWords(word, 7)))
	ShowModalInput()
}

func wordClickedHandler(word string) {
	SetModalTitle(word)
	if dicmgr.Lookup(word) {
		SetModalTitle(wordLinkHtml(word))
		go showWordDefinitionInModal(word)
	} else {
		showPossibleWords(word)
	}
	openModal()
}
