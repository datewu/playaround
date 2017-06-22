package main

import (
	"io"
	"log"
	"net"
	"os"
)

func forward(c net.Conn, port string) {
	client, err := net.Dial("tcp", port)
	if err != nil {
		log.Fatalf("Dial faild: %v", err)
	}
	log.Printf("Connected to localhost %v\n", c)
	go func() {
		defer client.Close()
		defer c.Close()
		io.Copy(client, c)
	}()

	go func() {
		defer client.Close()
		defer c.Close()
		io.Copy(c, client)
	}()

}
func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Usage %s listen:port forward:port \n", os.Args[0])
	}
	sport, dport := ":"+os.Args[1], ":"+os.Args[2]
	listener, err := net.Listen("tcp", sport)
	if err != nil {
		log.Fatalf("Failed to setup listener: %v", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("ERROR: failed to accept listener: %v", err)
		}
		log.Printf("Accepted connection %v\n", conn)
		go forward(conn, dport)

	}
}
