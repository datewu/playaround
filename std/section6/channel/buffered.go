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

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("main - starting...")

	done := make(chan *book, len(books))
	go buildHistograms(books, done)

	for b := range done {
		log.Printf("main -- Got %s", b.title)

	}

	log.Println("main - Done")
}

func buildHistograms(bs []*book, done chan<- *book) {

	for _, b := range books {

		log.Printf("Processing %s...", b.title)
		b.hist.chars = make(map[rune]int)

		file, err := os.Open(b.path)
		if err != nil {
			log.Println(err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		//		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			for _, c := range scanner.Text() {
				b.hist.chars[c]++
			}
		}

		log.Printf("Done with %s", b.title)
		done <- b
	}
	close(done)

}

func printHist(h *histogram) {
	for k, v := range h.chars {
		fmt.Printf("%q=%d, ", k, v)

	}
}
