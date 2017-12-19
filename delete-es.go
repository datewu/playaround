package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	client http.Client
	u      = flag.String("u", "http://es:9200/*-", "default es server addresss")
)

func main() {
	flag.Parse()

	var tries int
	days := -20

	go func() {
		http.ListenAndServe(":39999", nil)
	}()

	for {
		now := time.Now()
		indexDate := now.AddDate(0, 0, days).Format("2006.01.02")
		r, err := curl(*u + indexDate)
		if err != nil && tries < 3 {
			log.Println(tries, err)
			tries++
			continue
		}
		tries = 0
		fmt.Println(r)
		time.Sleep(24 * time.Hour)
	}
}

func curl(url string) (r string, err error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	r = string(body)

	return
}
