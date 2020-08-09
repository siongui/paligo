package main

import (
	"bytes"
	"html/template"

	. "github.com/siongui/godom"
)

const pwt = `
<style>
div.is-possible-word:hover {
  color: red;
  background-color: #F0F8FF;
}
</style>
<div class="columns">
  <div class="column is-narrow">
    <input class="input" type="text" value="{{.Word}}">
    {{range $possibleWord := .PossibleWords}}
      <div class="is-possible-word"
           onclick="pwh('{{$possibleWord}}')">
             {{$possibleWord}}
      </div>
    {{end}}
  </div>
  <div class="column" id="modalExp">Auto</div>
</div>
`

type pws struct {
	Word          string
	PossibleWords []string
}

func onPossibleWordHandler(word string) {
	Document.GetElementById("modalExp").SetInnerHTML("<div>" + wordLinkHtml(word) + "</div>")
}

func GetPossibleWordsHtml(word string, possibleWords []string) string {
	Document.Set("pwh", onPossibleWordHandler)

	pw := pws{word, possibleWords}
	t, err := template.New("pwt").Parse(pwt)
	if err != nil {
		return err.Error()
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, pw)
	if err != nil {
		return err.Error()
	}
	return buf.String()

}
