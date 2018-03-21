package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()
	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Response URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://non-exist-web.site.ya")
}
