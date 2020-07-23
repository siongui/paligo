package main

import (
	"bytes"
	"html/template"
	"net/http"

	"github.com/siongui/gopalilib/lib"
	sg "github.com/siongui/gopherjs-input-suggest"
)

var setmpl = `
{{range $bnwe := .}}
<article class="message">
  <div class="message-header">
    <p>{{$bnwe.BookName}}</p>
  </div>
  <div class="message-body">
    {{$bnwe.Explanation}}
  </div>
</article>
{{end}}`

func showWordPreviewByTemplate(wi lib.BookIdWordExps) {
	// bnwes: (Book-Name, Word-Explanation)s
	idexps := lib.BookIdWordExps2IdExpsAccordingToSetting(wi, bookIdAndInfos, getSetting(), navigatorLanguages)
	bnwes := lib.IdExps2BookNameWordExps(
		lib.ShortExplanation(idexps, bookIdAndInfos),
		bookIdAndInfos,
	)
	t1, _ := template.New("wordExplanationPreview").Parse(setmpl)
	// Google Search: go html template output string
	// https://groups.google.com/forum/#!topic/golang-nuts/dSFHCV-e6Nw
	var buf bytes.Buffer
	t1.Execute(&buf, bnwes)
	println(buf.String())
}

func httpGetWordJson2(w string) {
	resp, err := http.Get(HttpWordJsonPath(w))
	if err != nil {
		return
	}
	if resp.StatusCode != 200 {
		return
	}

	wi, err := lib.DecodeHttpRespWord(resp.Body)
	if err != nil {
		return
	}

	showWordPreviewByTemplate(wi)
}

func setupWordPreview() {
	sg.OnHighlightSelectedWord(func(word string) {
		//println(word)
		if !getSetting().IsShowWordPreview {
			return
		}
		//println("show word preview")
		go httpGetWordJson2(word)
	})
}
