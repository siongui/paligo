// +build ignore

package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/siongui/goef"
)

const gofile = `package {{.PkgName}}

import (
	"encoding/base64"
	"os"
	"path/filepath"
	"strings"
)

var virtualFilesystem = map[string]string{
{{ range .Files }}"{{ .Name }}": ` + "`" + `{{ .Base64Content }}` + "`" + `,
{{ end }}}

func ReadFile(filename string) ([]byte, error) {
	content, ok := virtualFilesystem[filename]
	if ok {
		if len(content) > 0 && content[0] == '#' {
			// this is a symlink
			p := filepath.Clean(filepath.Join(filepath.Dir(filename), content[1:]))
			return ReadFile(p)
		}
		if strings.HasSuffix(filename, ".html") || strings.HasSuffix(filename, ".js") || strings.HasSuffix(filename, ".css") || strings.HasSuffix(filename, ".json") {
			return []byte(content), nil
		}
		return base64.StdEncoding.DecodeString(content)
	}

	index, ok2 := virtualFilesystem["index.html"]
	if ok2 {
		return []byte(index), nil
	}

	return nil, os.ErrNotExist
}

func MapKeys() []string {
	keys := make([]string, len(virtualFilesystem))
	i := 0
	for k := range virtualFilesystem {
		keys[i] = k
		i++
	}
	return keys
}
`

func getFilenameContent(dirpath, path string, info os.FileInfo) (name, content string, err error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".js") || strings.HasSuffix(path, ".css") || strings.HasSuffix(path, ".json") {
		content = string(b)
		// escape backtick `
		content = strings.Replace(content, "`", "`"+`+"`+"`"+`"+`+"`", -1)
	} else {
		content = base64.StdEncoding.EncodeToString(b)
	}
	name, err = filepath.Rel(dirpath, path)
	return
}

func main() {
	websiteDir := flag.String("websiteDir", "", "output dir of website")
	flag.Parse()

	err := goef.GeneratePackage("main", *websiteDir, "offline/data.go", getFilenameContent, gofile)
	if err != nil {
		panic(err)
	}
}
