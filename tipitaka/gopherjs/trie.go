package main

import (
	"strconv"

	"github.com/siongui/go-succinct-data-structure-trie"
)

var ft bits.FrozenTrie

func Lookup(word string) bool {
	return ft.Lookup(word)
}

func GetSuggestedWords(word string, maxNum int) []string {
	return ft.GetSuggestedWords(word, maxNum)
}

func SetupTrie() {
	bits.SetAllowedCharacters("abcdeghijklmnoprstuvyāīūṁṃŋṇṅñṭḍḷ…'’° -")

	teDatab, err := ReadFile("trie-data.txt")
	if err != nil {
		panic(err)
	}
	teData := string(teDatab)
	rdDatab, err := ReadFile("trie-rank-directory-data.txt")
	if err != nil {
		panic(err)
	}
	rdData := string(rdDatab)
	nodeCountb, err := ReadFile("trie-node-count.txt")
	if err != nil {
		panic(err)
	}
	nodeCount, err := strconv.ParseUint(string(nodeCountb), 10, 64)
	if err != nil {
		panic(err)
	}
	ft = bits.FrozenTrie{}
	//println(teData)
	//println(rdData)
	//println(nodeCount)
	ft.Init(teData, rdData, uint(nodeCount))
}
