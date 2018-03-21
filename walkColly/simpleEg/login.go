package main

import (
	"log"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	err := c.Post("https://example.com/login", map[string]string{"username": "lol", "password": "pwd"})
	if err != nil {
		log.Fatal(err)
	}

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.Visit("https://example.com/")
}
