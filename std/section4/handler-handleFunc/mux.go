package main

import (
	"fmt"
	"log"
	"net/http"
)

// m is a custom request multiplexer
type m struct{}

func (m *m) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/greet":
		//	func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "hell0")
		//		}(w, r)

	default:
		http.NotFound(w, r)
	}
}

func mux() {
	log.Fatalln(http.ListenAndServe(":9090", &m{}))
}
