package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)
	conn, err := tls.Dial("tcp", ":4444", &tls.Config{InsecureSkipVerify: true})

	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	n, err := conn.Write([]byte("Hello ll\n"))
	if err != nil {
		log.Fatalln(n, err)
	}
	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Fatalln(n, err)
	}
	fmt.Println(string(buf[:n]))
}
