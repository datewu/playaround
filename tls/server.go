package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {
	log.SetFlags(log.Lshortfile)
	cer, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalln(err)
	}

	config := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", ":4444", config)
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go hf(conn)
	}
}
func hf(c net.Conn) {
	defer c.Close()
	r := bufio.NewReader(c)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(msg)
		n, err := c.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}

}
