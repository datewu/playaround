package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe("localhost:8090", nil)
}
