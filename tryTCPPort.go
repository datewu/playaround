package main

import "log"
import "net"
import "strings"
import "time"
import "fmt"

var retries = 5

func tcpAlive(port string) (res bool) {
	retries--
	log.Println(retries)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		if strings.Contains(err.Error(), "addreass already in use") {
			res = true
			if retries != 0 {
				time.Sleep(time.Second)
				return tcpAlive(port)
			}
			return

		}

	}
	l.Close()
	return
}

func main() {
	if tcpAlive("3000") {
		fmt.Println("service ok")
		return
	}
	fmt.Println("web not working")

}
