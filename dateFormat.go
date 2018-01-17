package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

const format = "Mon Jan 2 15:04:05 -0700 MST 2006"

func main() {
	t := time.Now()
	hash := t.Format("/06/01/02/") + strconv.Itoa(t.Nanosecond()/1000)
	hash2 := t.Format("/06/01/02/") + strconv.FormatInt(t.UnixNano(), 36)
	log.Println(t.Nanosecond())
	fmt.Println(hash, hash2)

	layout := "2006-01-02"
	dates := []string{
		"2017-12-01",
		"2018-11-10",
		"2018-01-03",
		"2018-03-15",
		"2019-01-01",
	}

	for _, s := range dates {
		tF, err := time.Parse(layout, s)
		if err != nil {
			log.Println(err)
		}
		h := tF.Format("/06/01/02/") + strconv.FormatInt(tF.UnixNano(), 36)
		fmt.Println(s, h)
	}
}
