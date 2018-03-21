package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.AllowedDomains = []string{"hackerspaces.org", "wiki.hackerspaces.org"}

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("something went wrong:", err)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visting", r.URL.String())
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

		// visit link found on page
		// only thoese links are visited which are in AllowedDomains
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// kick start
	c.Visit("https://hackerspaces.org/")
}
