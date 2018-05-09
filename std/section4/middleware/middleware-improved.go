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

func counter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		atomic.AddUint64(&requestCounter, 1)
		log.Println("COUNTER >> Counted")
	})
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

func use(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}

func main() {
	// http.Handle("/greet", logger(counter(greatHandler)))
	// http.Handle("/stats", logger(&statsHandler{}))
	h1 := use(http.HandlerFunc(greatHandler), counter, logger)
	h2 := use(&statsHandler{}, logger)
	http.Handle("/greet", h1)
	http.Handle("/stats", h2)

	log.Fatalln(http.ListenAndServe(":8080", nil))

}
