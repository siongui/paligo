package main

import (
	"regexp"
	"strings"

	. "github.com/siongui/godom"
)

var paliWord = regexp.MustCompile(`[AaBbCcDdEeGgHhIiJjKkLlMmNnOoPpRrSsTtUuVvYyĀāĪīŪūṀṁṂṃŊŋṆṇṄṅÑñṬṭḌḍḶḷ]+`)

func markPaliWordInSpan(s string) string {
	return paliWord.ReplaceAllStringFunc(s, func(match string) string {
		return "<span class='paliword'>" + match + "</span>"
	})
}

func toDom(s string, clickHandler func(string)) *Object {
	// wrap all words in span
	spanContainer := Document.CreateElement("span")
	spanContainer.SetInnerHTML(markPaliWordInSpan(s))

	// register click handler to every word
	spans := spanContainer.GetElementsByTagName("span")
	for _, span := range spans {
		word := strings.ToLower(span.InnerHTML())
		span.AddEventListener("click", func(e Event) {
			clickHandler(word)
		})
	}

	return spanContainer
}

// find all words in the element
func traverse(elm *Object, clickHandler func(string)) {
	// 1: element node
	if elm.NodeType() == 1 {
		for _, childNodes := range elm.ChildNodes() {
			traverse(childNodes, clickHandler)
		}
		return
	}

	// 3: text node
	if elm.NodeType() == 3 {
		s := elm.NodeValue()
		if strings.TrimSpace(s) != "" {
			// string is not whitespace
			elm.ParentNode().ReplaceChild(toDom(s, clickHandler), elm)
		}
		return
	}
}

func MarkEveryWord(selector string, clickHandler func(string)) {
	element := Document.QuerySelector(selector)
	traverse(element, clickHandler)
}
