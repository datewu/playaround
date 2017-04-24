package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"./links"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// Start with the command-line arguments.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unsean link.
	for i := 0; i < 2000; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()

			}
		}()

	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil && (strings.Contains(err.Error(), "dial") || strings.Contains(err.Error(), "too many")) {
		//log.Fatalln(err)
		log.Println(err)
		return nil
	}
	return list
}
