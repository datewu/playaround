package main

import "os"

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	// Add command-line arguments to worklist.
	go func() {
		worklist <- os.Args[1:]
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()

			}
		}()
	}

	// The main gorotine de-duplicates worklist items
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
