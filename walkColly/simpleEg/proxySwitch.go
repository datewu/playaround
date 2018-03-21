package main

import (
	"bytes"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate http and socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("http://127.0.0.1:1087", "socks5://192.168.0.137:1080")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	c.OnError(func(r *colly.Response, e error) {
		log.Println(r, e)
	})

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	// Fetch httpbin.org/ip five times
	for i := 0; i < 5; i++ {
		c.Visit("https://httpbin.org/ip")
	}
}
