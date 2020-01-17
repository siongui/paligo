package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
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

	total := 0
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
		total++
	}

	filenames := vfs.MapKeys()
	if len(filenames) == total {
		fmt.Println("total number of json file correct")
	} else {
		panic("total number of json files not correct")
	}
	for _, filename := range filenames {
		p := path.Join(wordJsonDir, filename)
		if _, err := os.Stat(p); err == nil {
			fmt.Println(p, "exist")
		} else {
			panic(p)
		}
	}
}
