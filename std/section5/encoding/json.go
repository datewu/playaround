package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	producer()
	consumer()
}

type proverb struct {
	ID       int
	Text     string
	reviewed bool
}

func producer() {
	ps := []proverb{
		proverb{ID: 1, Text: "Don't panic.", reviewed: true},
		proverb{ID: 2, Text: "Concurrency is not parallelism.", reviewed: true},
		proverb{ID: 3, Text: "Documentation is for users.", reviewed: true},
		proverb{ID: 4, Text: "The bigger the interface, the weaker the abstraction.", reviewed: true},
		proverb{ID: 5, Text: "Make the zero value useful.", reviewed: true},
	}

	fn := path.Join("..", "proverbs.json")
	f, err := os.Create(fn)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	if err := enc.Encode(ps); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("produces fineshed", fn)

}

func consumer() {
	fn := path.Join("..", "proverbs.json")
	f, err := os.Open(fn)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	var ps []proverb
	dec := json.NewDecoder(f)
	if err := dec.Decode(&ps); err != nil {
		log.Println(err)
		return
	}
	for _, p := range ps {
		fmt.Printf("%#v\n", p)
	}
}
