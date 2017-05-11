package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	i      int
	client http.Client
	u      = flag.String("u", "http://es:9200/*-", "default url")
)

func main() {
	flag.Parse()
	go func() {
		http.ListenAndServe(":39999", nil)
	}()

	for {
		now := time.Now()
		indexDate := now.AddDate(0, 0, -20).Format("2006.01.02")
		r, err := curl(*u + indexDate)
		if err != nil && i < 3 {
			log.Println(i, err)
			i++
			continue
		}
		i = 0
		log.Println(r)
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
	return string(body), nil
}
