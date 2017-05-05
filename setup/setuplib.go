package main

import (
	"encoding/json"
	"github.com/siongui/gopalilib/lib"
	"os"
)

func GetWordPath(word, wordsJsonDir string) string {
	return wordsJsonDir + "/" + word + ".json"
}

func GetBookIdWordExps(word, wordsJsonDir string) lib.BookIdWordExps {
	f, err := os.Open(GetWordPath(word, wordsJsonDir))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	w := lib.BookIdWordExps{}
	if err := dec.Decode(&w); err != nil {
		panic(err)
	}
	return w
}
