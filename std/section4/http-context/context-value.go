package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

type key int

const usernameKey key = 22

func addUsername(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := r.Header.Get("X-Username")
		if u == "" {
			u = "Annoymous"
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, usernameKey, u)
		h(w, r.WithContext(ctx))
	}
}

func check(ctx context.Context) {
	u := ctx.Value(usernameKey).(string)
	log.Printf("[%s] check performed\n", u)
}

func greet(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(usernameKey).(string)
	log.Printf("[%s] handling greet request\n", u)
	defer log.Printf("[%s] handled greet request\n", u)

	check(r.Context())
	w.Header().Set("X-Username", u)
	fmt.Fprintln(w, "Hello gopher")
}

func proverb(w http.ResponseWriter, r *http.Request) {
	u := r.Context().Value(usernameKey).(string)
	log.Printf("[%s] handling proverb request\n", u)
	defer log.Printf("[%s] handled proverb request\n", u)

	check(r.Context())
	w.Header().Set("X-Username", u)
	fmt.Fprintln(w, "Don't panic")
}

func main() {
	http.HandleFunc("/greet", addUsername(greet))
	http.HandleFunc("/proverb", addUsername(proverb))

	log.Println("starting server...")
	http.ListenAndServe(":8080", nil)
}
