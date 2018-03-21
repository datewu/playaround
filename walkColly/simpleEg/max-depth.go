package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		l := e.Attr("href")
		fmt.Println(l, "===>", e.Request.AbsoluteURL(l))

		e.Request.Visit(l)
	})

	c.Visit("https://en.wikipedia.org")
}
