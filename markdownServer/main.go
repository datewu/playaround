package main

import (
	"io/ioutil"
	"net/http"

	"github.com/russross/blackfriday"
)

func main() {

	http.HandleFunc("/", markdown)
	http.ListenAndServe(":8080", nil)
}

func markdown(w http.ResponseWriter, r *http.Request) {
	input, _ := ioutil.ReadFile("lol.md")
	output := blackfriday.MarkdownCommon(input)
	w.Header().Set("Content-Type", "text/html")
	w.Write(output)
}
