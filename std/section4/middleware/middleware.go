package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

var requestCounter uint64

func greatHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
	log.Println("GREETED")
}

type statsHandler struct{}

func (s *statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Request count: %d\n", atomic.LoadUint64(&requestCounter))
	log.Println("STATS PROVIDED")
}

func counter(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//h.ServeHTTP(w, r)
		h(w, r)
		atomic.AddUint64(&requestCounter, 1)
		log.Println("COUNTER >> Counted")
	}
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Println("LOGGER >> START", r.Method, r.URL.String())
			t := time.Now()
			next.ServeHTTP(w, r)
			log.Println("LOGGER >> end", r.Method, r.URL.String(), time.Now().Sub(t))
		})
}

func main() {
	http.Handle("/greet", logger(counter(greatHandler)))
	http.Handle("/stats", logger(&statsHandler{}))

	log.Fatalln(http.ListenAndServe(":8080", nil))

}

//
// Two styles of middleware because of the following:
// http.Handle("/", http.HandlerFunc(f)) equivalent to
// http.HandleFunc("/", f)
// if f has signature func(http.ResponseWriter, *http.Response)

func middlewareUsingHandlerFunc(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// the middlerware's logic here...
		f(w, r) // equivalent to f.ServeHTTP(w, r)
	}
}

func middlewareUsingHander(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// the middlerware's logic here...
		next.ServeHTTP(w, r)
	})
}
