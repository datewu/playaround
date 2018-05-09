package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
)

type book struct {
	title string
	path  string
}

type histogram struct {
	chars map[rune]int
}

type job struct {
	book *book
}

type result struct {
	job  job
	hist *histogram
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

var jobStream = make(chan job, 2)
var resultStream = make(chan result, 2)

func worker(g *sync.WaitGroup) {
	for j := range jobStream {
		r := result{j, buildHistogram(j.book)}
		resultStream <- r
	}
	g.Done()
}

func setupWorkerPool(size int) {
	var wg sync.WaitGroup
	wg.Add(size)
	for i := 0; i < size; i++ {
		go worker(&wg)
	}
	wg.Wait()
	close(resultStream)
}

func buildHistogram(b *book) *histogram {
	log.Println("buildStage processing-", b.title)
	h := histogram{chars: make(map[rune]int)}

	f, err := os.Open(b.path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		for _, c := range scanner.Text() {
			h.chars[c]++
		}
	}
	return &h
}

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)
	log.Println("Starting...")

	go func() {
		setupWorkerPool(10)
	}()

	go func() {
		for _, b := range books {
			jobStream <- job{b}
		}
		close(jobStream)
	}()
	for r := range resultStream {
		log.Printf("Job for %s done", r.job.book.title)
		printHist(r.hist)
	}

	log.Println("Done")
}

func printHist(h *histogram) {
	for k, v := range h.chars {
		fmt.Printf("%q=%d, ", k, v)

	}
}
