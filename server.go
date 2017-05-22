package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	log.Println(http.ListenAndServe(":445", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
