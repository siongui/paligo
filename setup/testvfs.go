package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"

	vfs "github.com/siongui/gopaliwordvfs"
)

const websiteDir = "../website"
const wordJsonDir = websiteDir + "/json/"

func main() {
	files, err := ioutil.ReadDir(wordJsonDir)
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		bVfs, err := vfs.ReadFile(file.Name())
		if err != nil {
			panic(err)
		}

		bReal, err := ioutil.ReadFile(path.Join(wordJsonDir, file.Name()))
		if err != nil {
			panic(err)
		}

		if !bytes.Equal(bVfs, bReal) {
			panic(file.Name())
		}

		fmt.Println(i, file.Name(), "ok")
	}
}
