package main

import (
	. "github.com/siongui/godom"
	dic "github.com/siongui/gopalilib/lib/dictionary"
)

func ActionXmlUrl(action string) string {
	return "https://siongui.github.io/tipitaka-romn/" + action
}

func GetXslUrl() string {
	return "https://siongui.github.io/tipitaka-romn/cscd/tipitaka-latn.xsl"
}

func HttpWordJsonPath(word string) string {
	if isOffline() {
		return "/json/" + word + ".json"
	}
	return "https://siongui.github.io/xemaauj9k5qn34x88m4h/" + word + ".json"
}

func isOffline() bool {
	return Window.Location().Hostname() == "localhost" && Window.Location().Port() == "8080"
}

func wordDictionaryUrl(word string) string {
	return "https://dictionary.sutta.org" + dic.WordUrlPath(word)
}
