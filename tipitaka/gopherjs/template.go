package main

import (
	"bytes"
	"html/template"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	"github.com/siongui/gopalilib/lib/dicmgr"
)

const pwt = `
<style>
.previewWordName {
  color: GoldenRod;
  font-weight: bold;
  font-size: 1.5rem;
  margin: .5rem;
}

div.is-possible-word:hover {
  color: red;
  background-color: #F0F8FF;
  cursor: pointer;
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
  <div class="column" id="modalExp">Click word to get definition</div>
</div>
`

const HtmlTemplateWordPreview = `
{{range $bnwe := .BookNameShortExps}}
<article class="message">
  <div class="message-header">
    <p>{{$bnwe.BookName}}</p>
  </div>
  <div class="message-body">
    {{$bnwe.Explanation}}
  </div>
</article>
{{end}}`

type pws struct {
	Word          string
	PossibleWords []string
}

func SetModalExp(html string) {
	Document.GetElementById("modalExp").SetInnerHTML(html)
}

func onPossibleWordHandler(word string) {
	SetModalExp("Loading " + wordLinkHtml(word) + " ...")

	go func() {
		wi, err := lib.HttpGetWordJson(HttpWordJsonPath(word))
		if err != nil {
			SetModalExp("Fail to Get " + word + ": " + err.Error())
			return
		}
		setting := lib.GetDefaultPaliSetting()

		html := `<div class="previewWordName">` + wordLinkHtml(word) + `</div>`
		//html += dicmgr.GetWordPreviewHtmlWithCustomTemplate(word, wi, setting, Window.Navigator().Languages(), HtmlTemplateWordPreview)
		html += dicmgr.GetWordDefinitionHtml(wi, setting, Window.Navigator().Languages())
		SetModalExp(html)
	}()
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
