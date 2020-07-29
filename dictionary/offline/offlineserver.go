package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Path
	if strings.HasSuffix(key, "/") {
		key += "index.html"
	}
	key = strings.TrimPrefix(key, "/")

	b, err := ReadFile(key)
	if err != nil {
		fmt.Fprintf(w, "404 not found")
		return
	}

	if strings.HasSuffix(key, ".js") {
		w.Header().Set("Content-Type", "text/javascript")
	}
	if strings.HasSuffix(key, ".css") {
		w.Header().Set("Content-Type", "text/css")
	}
	if strings.HasSuffix(key, ".json") {
		w.Header().Set("Content-Type", "application/json")
	}

	_, err = io.Copy(w, bytes.NewReader(b))
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
