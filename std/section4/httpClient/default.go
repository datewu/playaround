package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server >> Request received", r.Method, r.RequestURI)
		msg := "hello Gopher"
		log.Println("Server >> Sending", msg)
		time.Sleep(3 * time.Second)
		fmt.Fprintln(w, msg)
	}))
	defer ts.Close()

	log.Println("Client >> Making request ro test server", ts.URL)

	t := time.Now()
	r, err := http.Get(ts.URL)
	if err != nil {
		log.Fatalln(err)
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}
	msg := strings.TrimSpace(string(b))
	log.Printf("Client >> Received response %q in %v", msg, time.Since(t))
}
