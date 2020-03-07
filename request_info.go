package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	u := r.URL
	fmt.Println(r.Host, "u:", u, "schema:", u.Scheme, "path:", u.Path, "hostname():", u.Hostname(), u.String())
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
