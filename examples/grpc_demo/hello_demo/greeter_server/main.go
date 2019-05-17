package main

import (
	"context"
	"fmt"
	pb "github.com/aaronize/grandonion/examples/grpc_demo/hello_demo/helloworld"
	"google.golang.org/grpc"

	"log"
	"net"
)

const port = ":50051"

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	li, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		fmt.Printf("listening: %s\n", port)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(li); err != nil {
		log.Fatalf("failed to serve: %v", err)
	} else {
		fmt.Printf("sueecess serving...\n")
	}
}
