package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type book struct {
	title string
	path  string
	hist  histogram
}

type histogram struct {
	chars map[rune]int
}

var books = []*book{
	&book{title: "The Iliad", path: "../data/the-iliad.txt"},
	&book{title: "The Underground Railroad", path: "../data/the-underground-railroad.txt"},
	&book{title: "Pride and Prejudice", path: "../data/pride-and-prejudice.txt"},
	&book{title: "The Republic", path: "../data/the-republic.txt"},
	&book{title: "My Bondage and My Freedom", path: "../data/my-bondage-and-my-freedom.txt"},
	&book{title: "War and Peace", path: "../data/war-and-peace.txt"},
	&book{title: "Moby Dick", path: "../data/moby-dick.txt"},
	&book{title: "Meditations", path: "../data/meditations.txt"},
}

func collectStage(bs []*book) <-chan *book {
	out := make(chan *book)
	go func() {
		for _, b := range bs {
			log.Println("collectStage -", b.title)
			out <- b
		}
		close(out)
	}()
	return out
}

func buildStage(in <-chan *book) <-chan *book {
	out := make(chan *book)
	go func() {
		for b := range in {
			log.Println("buildStage processing-", b.title)
			b.hist.chars = make(map[rune]int)

			f, err := os.Open(b.path)
			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)

			for scanner.Scan() {
				for _, c := range scanner.Text() {
					b.hist.chars[c]++
				}
			}
			log.Println("buildStage Done-", b.title)
			out <- b
		}
		close(out)
	}()
	return out
}

func tallyStage(h *histogram, in <-chan *book) <-chan *book {
	out := make(chan *book)

	go func() {
		for b := range in {
			log.Printf("tallyStage - %s", b.title)
			for k, v := range b.hist.chars {
				h.chars[k] += v
			}
			out <- b
		}
		close(out)
	}()
	return out
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Starting...")

	hist := histogram{chars: make(map[rune]int)}

	cs := collectStage(books)
	bs := buildStage(cs)
	ts := tallyStage(&hist, bs)

	for b := range ts {
		log.Println("main -", b.title)
	}

	log.Println("main - HISTOGRAMd")
	printHist(&hist)
}

func printHist(h *histogram) {
	for k, v := range h.chars {
		fmt.Printf("%q=%d, ", k, v)

	}
}
