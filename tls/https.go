package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hw)
	err := http.ListenAndServeTLS(":3334", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}

func hw(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
