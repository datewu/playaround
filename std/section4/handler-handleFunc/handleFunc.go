package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleFunc() {
	http.HandleFunc("/", handler)
	log.Println("starting server...")
	log.Fatalln(http.ListenAndServe(":9090", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm ok")
}
