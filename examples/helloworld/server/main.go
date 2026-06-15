package main

import (
	"log"
	"net"

	pb "github.com/rafaelalmeidab/fc-grpc/examples/helloworld/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:8282")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}