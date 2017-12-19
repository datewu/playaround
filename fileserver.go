package main

import (
	"log"
	"net/http"
)

func main() {
	port := ":8090"
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Println("Listen on", port)
	http.ListenAndServe(port, nil)
}
