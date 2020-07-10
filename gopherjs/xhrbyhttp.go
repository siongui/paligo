package main

import (
	"net/http"

	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/lib"
)

func httpGetWordJson(w string, changeUrl bool) {
	resp, err := http.Get(HttpWordJsonPath(w))
	if err != nil {
		mainContent.Set("textContent", "Not Found")
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		mainContent.Set("textContent", "Not Found")
		return
	}

	wi := DecodeHttpRespWord(resp.Body)

	if changeUrl {
		Window.History().PushState(w, "", lib.WordUrlPath(w))
	}

	showWordByTemplate(wi)
}
