package main

import (
	vfs "github.com/siongui/gopaliwordvfs"
)

func main() {
	b, err := vfs.ReadFile("sacca.json")
	if err != nil {
		panic(err)
	}
	println(string(b))
}
