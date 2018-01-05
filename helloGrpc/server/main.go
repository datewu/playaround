package main

import (
	"log"
	"net"

	pb "../helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":8080"

// server is used to implement helloworld.GreeterServer.
type server bool

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Msg: "nihao: " + in.Name}, nil
}

func (s *server) Bye(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Msg: "zouni : " + in.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln("failed to listen", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, new(server))

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("faild to server", err)
	}

}
