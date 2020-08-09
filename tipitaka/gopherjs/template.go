package main

import (
	"bytes"
	"html/template"
)

const pwt = `
<input class="input" type="text" value="{{.Word}}">
{{range $possibleWord := .PossibleWords}}
<div>{{$possibleWord}}</div>
{{end}}
`

type pws struct {
	Word          string
	PossibleWords []string
}

func GetPossibleWordsHtml(word string, possibleWords []string) string {
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
