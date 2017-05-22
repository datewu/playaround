package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// tcpAlive return false if port available in 5 seconds
var retry = 5

func tcpAlive(port string) (res bool) {
	retry--
	log.Println(retry)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		if strings.Contains(err.Error(), "address already in use") {
			res = true
			if retry != 0 {
				time.Sleep(time.Second)
				return tcpAlive(port)
			}
		}
		return
	}
	l.Close()
	return
}

func main() {
	if tcpAlive("3000") {
		fmt.Println("service ok")
		return
	}
	fmt.Println("web not ok")
}
