package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"

	pb "../helloworld"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8080"
	name = "lol"
)

func main() {

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("did not connect", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	n := name
	if len(os.Args) > 1 {
		n = os.Args[1]
	}
	request := &pb.HelloRequest{n}
	r, err := c.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalln("could not greet:", err)
	}

	fmt.Println("Greeting", r.Msg)

	r, err = c.Bye(context.Background(), request)
	if err != nil {
		log.Fatalln("could not bye:", err)
	}

	fmt.Println("Bye", r.Msg)

}
