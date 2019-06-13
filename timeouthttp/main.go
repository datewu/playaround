package main

import (
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)
	w.WriteHeader(200)
	w.Write([]byte("200 - ok bad happened!"))
}

func main() {
	http.HandleFunc("/", greet)
	svc := &http.Server{
		Addr:         ":8090",
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}
	svc.ListenAndServe()
}
