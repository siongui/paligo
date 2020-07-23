package main

import (
	"bytes"
	"html/template"
	"net/http"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
	sg "github.com/siongui/gopherjs-input-suggest"
)

type WordPreview struct {
	Word              string
	BookNameShortExps []lib.BookNameWordExp
}

var setmpl = `
<span class="previewWordName">{{ .Word }}</span>
{{range $bnwe := .BookNameShortExps}}
<div class="shortDicExp">
  <span>{{$bnwe.BookName}}</span>
  <span>{{$bnwe.Explanation}}</span>
</div>
{{end}}`

var wordPreviewElm *Object

func setWordPreviewUI(word, rawhtml string) {
	wordPreviewElm.SetInnerHTML(rawhtml)
	wordPreviewElm.ClassList().Remove("is-hidden")
	Document.QuerySelector(".suggest").ClassList().Add("suggest-is-absolute")
	w := Document.QuerySelector(".suggest").Get("offsetWidth").String() + "px"
	//println(w)
	wordPreviewElm.Style().SetLeft(w)
}

func showWordPreviewByTemplate(word string, wi lib.BookIdWordExps) {
	// bnwes: (Book-Name, Word-Explanation)s
	idexps := lib.BookIdWordExps2IdExpsAccordingToSetting(wi, bookIdAndInfos, getSetting(), navigatorLanguages)
	bnwes := lib.IdExps2BookNameWordExps(
		lib.ShortExplanation(idexps, bookIdAndInfos),
		bookIdAndInfos,
	)
	t1, _ := template.New("wordExplanationPreview").Parse(setmpl)
	wp := WordPreview{word, bnwes}
	// Google Search: go html template output string
	// https://groups.google.com/forum/#!topic/golang-nuts/dSFHCV-e6Nw
	var buf bytes.Buffer
	t1.Execute(&buf, wp)
	rawhtml := buf.String()

	setWordPreviewUI(word, rawhtml)
}

func httpGetWordJson2(word string) {
	resp, err := http.Get(HttpWordJsonPath(word))
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

	showWordPreviewByTemplate(word, wi)
}

func setupWordPreview() {
	wordPreviewElm = Document.QuerySelector(".suggestedWordPreview")
	sg.OnHighlightSelectedWord(func(word string) {
		//println(word)
		if !getSetting().IsShowWordPreview {
			return
		}
		//println("show word preview")
		go httpGetWordJson2(word)
	})

	sg.OnUpdateSuggestMenu(func(word string) {
		wordPreviewElm.ClassList().Add("is-hidden")
		Document.QuerySelector(".suggest").ClassList().Remove("suggest-is-absolute")
	})

	sg.OnHideSuggestMenu(func() {
		wordPreviewElm.ClassList().Add("is-hidden")
		Document.QuerySelector(".suggest").ClassList().Remove("suggest-is-absolute")
	})
}
