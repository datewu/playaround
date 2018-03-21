package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gocolly/colly"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	c := colly.NewCollector(
		colly.UserAgent("xxx"),
	)
	c.AllowURLRevisit = true

	c.OnRequest(func(r *colly.Request) {
		r.UserAgent = randomString()
		fmt.Println("visting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("visted", r.URL)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnXML("//h1", func(e *colly.XMLElement) {
		fmt.Println(e.Text)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.URL)
	})

	fmt.Println("vim-go")
}

func randomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
