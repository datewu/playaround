package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type greetHandler struct{}

func (g *greetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello gopher!")
}

type proverb struct {
	id    int
	value string
}

type proverbsHandler struct {
	proverbs []proverb
}

func newProverbHandler() *proverbsHandler {
	return &proverbsHandler{
		[]proverb{
			proverb{id: 1, value: "Don't panic."},
			proverb{id: 2, value: "Concurrency is not parallelism."},
			proverb{id: 3, value: "Documentation is for users."},
			proverb{id: 4, value: "The bigger the interface, the weaker the abstraction."},
			proverb{id: 5, value: "Make the zero value useful."},
		},
	}
}

func (p *proverbsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/proverbs/"):])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	proverb, err := p.lookup(id)
	if err == errUnknownProverb {
		http.Error(w, errUnknownProverb.Error(), http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, proverb.value)
}

var errUnknownProverb = errors.New("unknown proverb")

func (p *proverbsHandler) lookup(id int) (*proverb, error) {
	for _, proverb := range p.proverbs {
		if id == proverb.id {
			return &proverb, nil
		}
	}
	return nil, errUnknownProverb
}

func main() {
	g := new(greetHandler)
	http.Handle("/great/", g)

	p := newProverbHandler()
	http.Handle("/proverbs/", p)

	http.HandleFunc("/", handler)
	log.Println("starting server...")
	log.Fatalln(http.ListenAndServe(":9090", nil))
}
