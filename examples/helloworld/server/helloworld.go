package main

import (
	"context"
	"log"

	pb "github.com/rafaelalmeidab/fc-grpc/examples/helloworld/proto"
)

func (s *Server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Greet call received")

	return &pb.HelloReply{Message: "Hello " + request.Name}, nil
}