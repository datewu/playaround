package main

import (
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/go-playground/lars"
	// gp "github.com/gorilla/http"
)

var tr *http.Transport
var cl *http.Client

var sema = make(chan struct{}, 200)

func main() {

	tr = &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   5,
	}
	cl = &http.Client{Transport: tr}
	handler := lars.New()
	handler.Get("/", ShowIndex)

	go func() {

		if err := http.ListenAndServe(":9898", handler.Serve()); err != nil {
			fmt.Println("start server error:", err)
		}
	}()

	h := lars.New()
	h.Get("/test", ShowHello)
	if err := http.ListenAndServe(":9899", h.Serve()); err != nil {
		fmt.Println("start server error:", err)
	}
}

func ShowIndex(c lars.Context) {
	//client := &http.Client{Transport: tr}
	sema <- struct{}{}
	req, err := http.NewRequest("GET", "http://localhost:9899/test", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := cl.Do(req)
	if err != nil {
		fmt.Println(err)
		c.Response().Write([]byte("err"))
		return
	}
	defer resp.Body.Close()
	time.Sleep(5 * time.Second)
	c.Response().Write([]byte(resp.Status))
	<-sema
	//	gp.Get(c.Response(), "http://localhost:9899/test")
}

func ShowHello(c lars.Context) {
	c.Response().Write([]byte("ok"))
}
