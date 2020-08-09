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
/* for suggestion preview */
.suggestedWordPreview {
  border-top-color: #C9D7F1;
  border-right-color: #36C;
  border-bottom-color: #36C;
  border-left-color: #A2BAE7;
  border-style: solid;
  border-width: 1px;
  z-index: 10;
  padding: 0;
  background-color: white;
  overflow: hidden;
  position: absolute;
  text-align: left;
  font-size: large;
  border-radius: 4px;
  margin-top: 1px;
  line-height: 1.25em;
//  width: 32em;
//  text-align: left;
  /* http://stackoverflow.com/questions/12128465/twitter-bootstrap-break-word-not-work-on-dropdown-menu */
  word-wrap: break-word;
  white-space: normal;
}

.previewWordName {
  color: GoldenRod;
  font-weight: bold;
  font-size: 1.5em;
  margin: .5em;
}

div.shortDicExp:hover {
  font-size: 1.5em;
  line-height: 1em;
  background-color: #F0F8FF;
  border: 1px dotted aqua;
}

div.shortDicExp > span:first-child {
  color: red;
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
		SetModalExp(dicmgr.GetWordPreviewHtml(word, wi, setting, Window.Navigator().Languages()))
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
