package main

func HttpWordJsonPath(word string) string {
	if isOffline() {
		return "/json/" + word + ".json"
	}
	return "https://siongui.github.io/xemaauj9k5qn34x88m4h/" + word + ".json"
	//return "/xemaauj9k5qn34x88m4h/" + word + ".json"
}
