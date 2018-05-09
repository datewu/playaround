package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"time"
)

func main() {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(3 * time.Second)
			fmt.Fprintln(w, "lol")
		}))
	defer ts.Close()
	d := net.Dialer{
		//Timeout: time.Millisecond * 10,
		Timeout: time.Nanosecond * 10,
	}

	c := http.Client{
		Transport: &http.Transport{
			Dial: d.Dial,
		},
	}

	t := time.Now()
	_, err := c.Get(ts.URL)
	if err != nil {
		log.Println(err)
	}
	log.Println("time elasped", time.Now().Sub(t))
}
