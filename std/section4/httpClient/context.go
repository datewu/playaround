package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server >> Request received", r.Method, r.RequestURI)
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		time.Sleep(time.Duration(rnd.Intn(8)) * time.Second)
		msg := "hello Gopher"
		log.Println("Server >> Sending", msg)
		fmt.Fprintln(w, msg)
	}))
	defer ts.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, _ := http.NewRequest(http.MethodGet, ts.URL, nil)
	req = req.WithContext(ctx)

	t := time.Now()
	log.Println("Sending request...")

	log.Println("Client >> Making request ro test server", ts.URL)
	r, err := http.DefaultClient.Do(req)
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
