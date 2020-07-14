package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func SendRootIndex(w io.Writer) {
	f, err := os.Open("website/index.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer f.Close()

	_, err = io.Copy(w, f)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "my 404 page!")
	SendRootIndex(w)
}

func FileServerWithCustom404(fs http.FileSystem) http.Handler {
	fsh := http.FileServer(fs)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fs.Open(path.Clean(r.URL.Path))
		if os.IsNotExist(err) {
			NotFound(w, r)
			return
		}
		fsh.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":8000", FileServerWithCustom404(http.Dir("website")))
}
